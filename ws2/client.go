/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-22 21:48:55
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-23 15:43:57
 */
package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kataras/iris/v12/websocket"
)

const (
	endpoint = "ws://127.0.0.1:8000/ws"
	// endpoint              = "ws://127.0.0.1:8000/ws"
	namespace             = "default"
	dialAndConnectTimeout = 5 * time.Second
)

// this can be shared with the server.go's.
// `NSConn.Conn` has the `IsClient() bool` method which can be used to
// check if that's is a client or a server-side callback.
var clientEvents = websocket.Namespaces{
	namespace: websocket.Events{
		websocket.OnNamespaceConnected: func(c *websocket.NSConn, msg websocket.Message) error {
			log.Printf("connected to namespace: %s", msg.Namespace)
			return nil
		},
		websocket.OnNamespaceDisconnect: func(c *websocket.NSConn, msg websocket.Message) error {
			log.Printf("disconnected from namespace: %s", msg.Namespace)
			return nil
		},
		"chat": func(c *websocket.NSConn, msg websocket.Message) error {
			// log.Printf("%s", string(msg.Body))
			// c.Emit("chat", []byte("------->"))
			fmt.Printf("\n client recv --------------------- [%s]:[%s]", c.Conn.ID(), string(msg.Body))
			return nil
		},
	},
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(dialAndConnectTimeout))
	defer cancel()

	// username := "my_username"
	// dialer := websocket.GobwasDialer(websocket.GobwasDialerOptions{Header: websocket.GobwasHeader{"X-Username": []string{username}}})
	dialer := websocket.DefaultGobwasDialer
	client, err := websocket.Dial(ctx, dialer, endpoint, clientEvents)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	c, err := client.Connect(ctx, namespace)
	if err != nil {
		panic(err)
	}

	c.Emit("chat", []byte("Hello from Go client side!"))

	for {
		time.Sleep(time.Second * 5)
		text := []byte(time.Now().String())

		if bytes.Equal(text, []byte("exit")) {
			if err := c.Disconnect(nil); err != nil {
				log.Printf("reply from server: %v", err)
			}
			break
		}

		ok := c.Emit("chat", text)
		if !ok {
			break
		}
	}
} // try running this program twice or/and run the server's http://localhost:8080 to check the browser client as well.
