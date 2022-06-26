/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-22 00:16:25
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-22 09:45:39
 */
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

const (
	enableJWT             = false
	endpoint              = "ws://localhost:9999/ws"
	namespace             = "default"
	dialAndConnectTimeout = 5 * time.Second
)

// if namespace is empty then simply websocket.Events{...} can be used instead.
var serverEvents = websocket.Namespaces{
	namespace: websocket.Events{
		websocket.OnNamespaceConnected: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// with `websocket.GetContext` you can retrieve the Iris' `Context`.
			ctx := websocket.GetContext(nsConn.Conn)

			log.Printf("[%s] connected to namespace [%s] with IP [%s]",
				nsConn, msg.Namespace,
				ctx.RemoteAddr())
			return nil
		},
		websocket.OnNamespaceDisconnect: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("[%s] disconnected from namespace [%s]", nsConn, msg.Namespace)
			return nil
		},
		"chat": func(nsConn *websocket.NSConn, msg websocket.Message) error {
			// room.String() returns -> NSConn.String() returns -> Conn.String() returns -> Conn.ID()
			log.Printf("[%s] sent: %s", nsConn, string(msg.Body))

			// Write message back to the client message owner with:
			// nsConn.Emit("chat", msg)
			// Write message to all except this client with:
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			return nil
		},
	},
}

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
			log.Printf(" client ----------> %s", string(msg.Body))
			return nil
		},
	},
}

func HelloClient(ctx iris.Context) {
	go Clinet()
	ctx.JSON(iris.Map{"hello": "world !"})
}

func main() {
	app := iris.New()
	websocketServer := websocket.New(
		websocket.DefaultGorillaUpgrader, /* DefaultGobwasUpgrader can be used too. */
		serverEvents)
	app.Any("/ws", websocket.Handler(websocketServer))
	app.Any("/", HelloClient)
	app.Listen(":9000")
}

func Clinet() {
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
	for i := 0; i < 10; i++ {
		fmt.Println("------------")
		c.Emit("chat", []byte("Hello from Go client side: "+time.Now().Format("2006-01-02 15:04:05")))
		time.Sleep(time.Second * 5)
	}
	return
}
