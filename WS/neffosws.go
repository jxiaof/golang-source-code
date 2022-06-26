/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-21 19:39:06
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-22 00:39:31
 */
package main

import (
	"fmt"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
)

func main() {
	fmt.Println("Hello World")
	app := iris.New()
	ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
		websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
			log.Printf("Server got: %s from [%s]", msg.Body, nsConn.Conn.ID())
			nsConn.Conn.Server().Broadcast(nsConn, msg)
			// if !nsConn.Conn.IsClient() {
			// 	fmt.Println("server")
			// 	// wsMsg := websocket.Message{

			// 	// 	Body: []byte("Hello from the server"),
			// 	// }
			// 	// nsConn.Conn.
			// }
			return nil
		},
	})

	ws.OnConnect = func(c *websocket.Conn) error {
		log.Printf("[%s] Connected to server!", c.ID())
		return nil
	}

	ws.OnDisconnect = func(c *websocket.Conn) {
		log.Printf("[%s] Disconnected from server", c.ID())
	}
	app.Any("/ws", websocket.Handler(ws))
	app.Any("/", func(ctx iris.Context) {
		ctx.Writef("Hello from the server!")
	})
	app.Listen(":9999")

}
