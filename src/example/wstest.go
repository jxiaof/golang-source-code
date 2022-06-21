package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/remotecommand"
)

// https://github.com/kubernetes/dashboard/blob/master/src/app/backend/handler/terminal.go

type PtyHandler interface {
	io.Reader
	io.Writer
	remotecommand.TerminalSizeQueue
}

const END_OF_TRANSMISSION = "\u0004"

// TerminalMessage is the messaging protocol between ShellController and TerminalSession.
//
// OP      DIRECTION  FIELD(S) USED  DESCRIPTION
// ---------------------------------------------------------------------
// bind    fe->be     SessionID      Id sent back from TerminalResponse
// stdin   fe->be     Data           Keystrokes/paste buffer
// resize  fe->be     Rows, Cols     New terminal size
// stdout  be->fe     Data           Output from the process
// toast   be->fe     Data           OOB message to be shown to the user
type TerminalMessage struct {
	Op, Data, SessionID string
	Rows, Cols          uint16
}

// TerminalSession
type TerminalSession struct {
	ID       string
	wsConn   *websocket.Conn
	sizeChan chan remotecommand.TerminalSize
	doneChan chan struct{}
}

// TerminalSize handles pty->process resize events
// Called in a loop from remotecommand as long as the process is running
func (t *TerminalSession) Next() *remotecommand.TerminalSize {
	select {
	case size := <-t.sizeChan:
		return &size
	case <-t.doneChan:
		return nil
	}
}

// Read handles pty->process messages (stdin, resize)
// Called in a loop from remotecommand as long as the process is running
func (t *TerminalSession) Read(p []byte) (int, error) {
	_, message, err := t.wsConn.ReadMessage()
	if err != nil {
		log.Printf("%s: read ws message failed: %v", t.ID, err)
		return copy(p, END_OF_TRANSMISSION), err
	}
	log.Printf("%s: read", t.ID)
	// TODO: msg type
	return copy(p, message), nil
	//var msg TerminalMessage
	//if err := json.Unmarshal(message, &msg); err != nil {
	//	log.Printf("json decoded failed: %v", err)
	//	return copy(p, END_OF_TRANSMISSION), err
	//}
	//switch msg.Op {
	//case "stdin":
	//	return copy(p, msg.Data), nil
	//case "resize":
	//	t.sizeChan <- remotecommand.TerminalSize{Width: msg.Cols, Height: msg.Rows}
	//	return 0, nil
	//default:
	//	log.Printf("unknown message type '%s'", msg.Op)
	//	return copy(p, END_OF_TRANSMISSION), fmt.Errorf("unknown message type '%s'", msg.Op)
	//}
}

// Write handles process->pty stdout
// Called from remotecommand whenever there is any output
func (t *TerminalSession) Write(p []byte) (int, error) {
	//msg, err := json.Marshal(TerminalMessage{
	//	Op:   "stdout",
	//	Data: string(p),
	//})
	//if err != nil {
	//	log.Printf("json encode failed: %v", err)
	//	return 0, err
	//}
	// TODO: msg type
	if err := t.wsConn.WriteMessage(websocket.TextMessage, p); err != nil {
		log.Printf("%s: write ws message failed: %v", t.ID, err)
		return 0, err
	}
	log.Printf("%s: write", t.ID)
	return len(p), nil
}

func (t *TerminalSession) Close() error {
	close(t.doneChan)
	return t.wsConn.Close()
}

func newTerminalSession(id string, r *http.Request, w http.ResponseWriter) (*TerminalSession, error) {
	conn, err := upGrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return &TerminalSession{
		ID:       id,
		wsConn:   conn,
		sizeChan: make(chan remotecommand.TerminalSize),
		doneChan: make(chan struct{}),
	}, nil
}

func ContainerExec(ctx context.Context, r *http.Request, w http.ResponseWriter, cluster, namespace, podName, container string) error {
	kclient, err := k8s.GetKubeClient(cluster)
	if err != nil {
		log.Println(err)
		return err
	}

	// 获取kube config配置
	config, err := k8s.GetClientConfig(cluster)
	if err != nil {
		log.Println(err)
		return err
	}

	cmd := []string{"sh", "-c", "test -f /bin/bash && bash || sh"}
	option := &corev1.PodExecOptions{
		Command:   cmd,
		Stdin:     true,
		Stdout:    true,
		Stderr:    true,
		TTY:       true,
		Container: container,
	}
	//ctx, cancel := context.WithTimeout(ctx, models.READ_LOG_TIMEOUT)
	//defer cancel()
	req := kclient.CoreV1().RESTClient().
		Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		VersionedParams(option, scheme.ParameterCodec).
		Timeout(models.READ_LOG_TIMEOUT)

	executor, err := remotecommand.NewSPDYExecutor(config, http.MethodPost, req.URL())
	if err != nil {
		log.Println(err)
		return err
	}

	sessID := fmt.Sprintf("%s|%s|%s|%s|%v", cluster, namespace, podName, container, time.Now().Unix())
	t, err := newTerminalSession(sessID, r, w)
	if err != nil {
		log.Println(err)
		return err
	}
	defer t.Close()

	if err = executor.Stream(remotecommand.StreamOptions{
		Stdin:             t,
		Stdout:            t,
		Stderr:            t,
		TerminalSizeQueue: t,
		Tty:               true,
	}); err != nil {
		// http连接已经被hijacked，http.ResponseWriter不能再使用，所以不返回err
		log.Printf("exec stream failed: %v", err)
		return nil
	}

	return nil
}
