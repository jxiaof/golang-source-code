package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/kataras/neffos"
	"github.com/kataras/neffos/gobwas"
	"github.com/kataras/neffos/gorilla"
)

const (
	endpoint  = "localhost:9090"
	namespace = "default"
	timeout   = 1800 * time.Second
)

type Users struct {
	mu      sync.RWMutex
	entries map[string]*userConn // key = user's unique identifier, i.e "username".
}

// returns true if new conn.
func (u *Users) conn(c *neffos.NSConn) (*userConn, bool) {
	user := c.Conn.ID()
	u.mu.RLock()
	entry, ok := u.entries[user]
	u.mu.RUnlock()

	if !ok {
		entry = &userConn{
			conns: make(map[*neffos.NSConn]struct{}),
		}

		u.mu.Lock()
		u.entries[user] = entry
		u.mu.Unlock()
	}

	entry.addConn(c)
	return entry, !ok
}

func (u *Users) get(c *neffos.NSConn) *userConn {
	u.mu.RLock()
	entry, ok := u.entries[c.Conn.ID()]
	u.mu.RUnlock()

	if !ok {
		return nil
	}

	return entry
}

func (u *Users) remove(user string) {
	u.mu.Lock()
	delete(u.entries, user)
	u.mu.Unlock()
}

type userConn struct {
	mu    sync.RWMutex
	conns map[*neffos.NSConn]struct{}
}

// returns true for new conn.
func (u *userConn) addConn(c *neffos.NSConn) bool {
	u.mu.RLock()
	_, ok := u.conns[c]
	u.mu.RUnlock()
	if !ok {
		u.mu.Lock()
		u.conns[c] = struct{}{}
		u.mu.Unlock()
		return true
	}

	return false
}

func (u *userConn) deleteConn(c *neffos.NSConn) (wasLast bool) {
	u.mu.Lock()
	delete(u.conns, c)
	fmt.Println("deleteConn:", len(u.conns))
	wasLast = len(u.conns) == 0
	u.mu.Unlock()

	return
}

func (u *userConn) Emit(event string, data []byte) (ok bool) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	for c := range u.conns {
		ok = c.Emit(event, data)
		if !ok {
			delete(u.conns, c)
		}
	}
	return
}

func (u *userConn) Disconnect(ctx context.Context) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	for c := range u.conns {
		c.Disconnect(ctx)
	}
}

func (u *userConn) Close() {
	u.mu.Lock()
	defer u.mu.Unlock()

	for c := range u.conns {
		c.Conn.Close()
		delete(u.conns, c)
	}
}

var users = &Users{
	entries: make(map[string]*userConn),
}

var handler = neffos.WithTimeout{
	ReadTimeout:  timeout,
	WriteTimeout: timeout,
	Namespaces: neffos.Namespaces{
		"default": neffos.Events{
			neffos.OnNamespaceConnected: func(c *neffos.NSConn, msg neffos.Message) error {
				_, isNew := users.conn(c)
				if isNew || c.Conn.IsClient() {
					log.Printf("[%s] connected to [%s].", c.Conn.ID(), msg.Namespace)
				}

				if !c.Conn.IsClient() {
					// c.Emit("chat", []byte("welcome to server's namespace"))
					s := "[ " + c.Conn.ID() + " ] welcome to server's namespace"
					users.get(c).Emit("chat", []byte(s))
				}
				return nil
			},
			neffos.OnNamespaceDisconnect: func(c *neffos.NSConn, msg neffos.Message) error {
				if msg.Err != nil {
					log.Printf("This client can't disconnect yet, server does not allow that action, reason: %v", msg.Err)
					return nil
				}

				conn := users.get(c)
				if conn == nil {
					return nil
				}

				wasLast := conn.deleteConn(c)

				if wasLast {
					users.remove(c.Conn.ID())
					log.Printf("[%s] disconnected from [%s].", c.Conn.ID(), msg.Namespace)
				}

				if c.Conn.IsClient() {
					os.Exit(0)
				}

				return nil
			},
			"chat": func(c *neffos.NSConn, msg neffos.Message) error {
				fmt.Println("chat entries:", len(users.entries))
				if len(msg.Body) < 1 {
					log.Printf("-")
				}
				// log.Printf("get msg ---------------------> [%s] %s", c.Conn.ID(), msg.Body)
				conn := users.entries[c.Conn.ID()]
				if conn == nil {
					fmt.Println("conn is nil")
					return nil
				}
				for k, _ := range conn.conns {
					// fmt.Println("------> k == c: ", k == c, v, c.Conn.IsClosed(), c.Conn.IsClient(), string(msg.Body))
					if k == c {
						break
					}

				}

				if !c.Conn.IsClient() {
					// this is possible too:
					// if bytes.Equal(msg.Body, []byte("force disconnect")) {
					// 	println("force disconnect")
					// 	return c.Disconnect()
					// }

					// log.Printf("--server-side-- send back the message [%s:%s]", msg.Event, string(msg.Body))
					// c.Emit(msg.Event, msg.Body)
					//	c.Server().Broadcast(nil, msg) // to all including this connection.
					// c.Server().Broadcast(c, msg) // to all except this connection.

					// users.get(c).Emit(msg.Event, msg.Body)
					// fmt.Println("------>", len(users.entries))
					s := time.Now().Format("2006-01-02 15:04:05") + fmt.Sprintf(" [ %s ] server send ------------>   ", c.Conn.ID()) + string(msg.Body)
					msg.Body = []byte(s)
					ok := users.get(c).Emit(msg.Event, msg.Body) // c.Emit(msg.Event, msg.Body)
					if !ok {
						// break
						return fmt.Errorf("failed to emit message to the client")
					}
					// users.get(c).Emit(msg.Event, msg.Body)
				}

				return nil
			},
		},
	},
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("expected program to start with 'server' or 'client' argument")
	}
	side := args[0]

	var (
		upgrader = gobwas.DefaultUpgrader
		dialer   = gobwas.DefaultDialer
	)

	if len(args) > 1 {
		method := args[1]
		if method == "gorilla" {
			upgrader = gorilla.DefaultUpgrader
			dialer = gorilla.DefaultDialer
			if side == "server" {
				log.Printf("Using with Gorilla Upgrader.")
			} else {
				log.Printf("Using with Gorilla Dialer.")
			}
		}
	}

	switch side {
	case "server":
		server(upgrader)
	case "client":
		client(dialer)
	default:
		log.Fatalf("unexpected argument, expected 'server' or 'client' but got '%s'", side)
	}
}

func server(upgrader neffos.Upgrader) {
	srv := neffos.New(upgrader, handler)
	srv.IDGenerator = func(w http.ResponseWriter, r *http.Request) string {
		// r.Header.Get("x-user-name")
		// return r.RemoteAddr[:strings.IndexByte(r.RemoteAddr, ':')]
		uid, err := uuid.NewRandom()
		if err != nil {
			log.Printf("failed to generate uid: %v", err)
			time.Sleep(time.Second)
			return time.Now().Format("2006-01-02 15:04:05")
		}
		return uid.String()
	}

	srv.OnConnect = func(c *neffos.Conn) error {
		log.Printf("[%s] connected to server. entrys : %d", c.ID(), len(users.entries))
		// time.Sleep(3 * time.Second)
		// c.Connect(nil, namespace) // auto-connect to a specific namespace.
		// c.Write(namespace, "chat", []byte("Welcome to the server (after namespace connect)"))
		// println("client connected")
		return nil
	}
	srv.OnDisconnect = func(c *neffos.Conn) {
		log.Printf("[%s] disconnected from the server.", c.ID())
	}
	srv.OnUpgradeError = func(err error) {
		log.Printf("ERROR: %v", err)
	}

	log.Printf("Listening on: %s\nPress CTRL/CMD+C to interrupt.", endpoint)
	// go func() {
	// 	fmt.Println("goroutineNum: ", runtime.NumGoroutine())
	// 	time.Sleep(3 * time.Second)
	// }()
	go http.ListenAndServe(endpoint, srv)

	// fmt.Fprint(os.Stdout, ">> ")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			log.Printf("ERROR: %v", scanner.Err())
			return
		}

		text := scanner.Bytes()
		if bytes.Equal(text, []byte("force disconnect")) {
			// for _, conn := range srv.GetConnectionsByNamespace(namespace) {
			// 	conn.Disconnect()
			// }
			// srv.Broadcast(nil, neffos.Message{
			// 	Namespace: namespace,
			// 	Event:     neffos.OnNamespaceDisconnect,
			// })
			srv.Do(func(c *neffos.Conn) {
				// c.Close()
				c.Namespace(namespace).Disconnect(nil)
			}, false)
		} else {
			// srv.Do(func(c *neffos.Conn) {
			// 	// c.Write(text)
			// 	c.Write(namespace, "chat", text)
			// }, false)
			srv.Broadcast(nil, neffos.Message{Namespace: namespace, Event: "chat", Body: text})
		}
		// fmt.Fprint(os.Stdout, ">> ")

	}
}

func client(dialer neffos.Dialer) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(timeout))
	defer cancel()

	client, err := neffos.Dial(ctx, dialer, endpoint, handler)
	if err != nil {
		panic(err)
	}

	defer client.Close()

	connectNamespaceTimeout, cancel2 := context.WithTimeout(context.Background(), timeout/2)
	defer cancel2()

	// c, err := client.WaitServerConnect(nil, namespace)
	c, err := client.Connect(connectNamespaceTimeout, namespace)
	if err != nil {
		panic(err)
	}
	// println("connected.")

	// fmt.Fprint(os.Stdout, ">> ")
	scanner := bufio.NewScanner(os.Stdin)
	// go func() {
	// 	t := time.NewTicker(3 * time.Second)
	// 	defer t.Stop()
	// 	select {
	// 	case <-t.C:
	// 		ok := c.Emit("chat", []byte(""))
	// 		if !ok {
	// 			fmt.Println("send msg err: %v", ok)
	// 			return
	// 		}
	// 		fmt.Println("goroutineNum: ", runtime.NumGoroutine())
	// 	case <-time.After(1800 * time.Second):
	// 		fmt.Println("--------- -- -- -- timeout -- -- -- ---------")
	// 		fmt.Println("goroutineNum: ", runtime.NumGoroutine())
	// 		return
	// 	}
	// }()
	for {
		if !scanner.Scan() {
			log.Printf("ERROR: %v", scanner.Err())
			return
		}

		text := scanner.Bytes()

		if bytes.Equal(text, []byte("exit")) {
			if err := c.Disconnect(nil); err != nil {
				// log.Printf("from server: %v", err)
			}
			continue
		}
		// fmt.Println("----->goroutineNum: ", runtime.NumGoroutine())
		ok := c.Emit("chat", text)
		if !ok {
			fmt.Println("send msg err: %v", ok)
			break
		}

		// fmt.Fprint(os.Stdout, ">> ")
	}

}
