/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-23 16:31:14
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-26 10:38:01
 */
package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gobwas"
)

var IdGen = func(ctx iris.Context) string {
	if ssid := ctx.GetHeader("X-User-Session"); ssid != "" {
		return ssid
	}
	return websocket.DefaultIDGenerator(ctx)
}

var RawEvent = websocket.Events{
	websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {
		// return DoSomething(nsConn, msg)
		id := nsConn.Conn.ID()
		log.Printf("Server got: %s from [%s]", msg.Body, id)

		msgStr := string(msg.Body)
		if len(msgStr) > 36 {
			li := bytes.Split(msg.Body, []byte(":"))
			log.Printf("client got id: %s", string(li[0]))
			log.Printf("client got msg: %s", string(li[1]))
			li2 := [][]byte{li[0], []byte("server"), li[1]}
			b := bytes.Join(li2, []byte(":"))
			nsConn.Conn.Socket().WriteText(b, 0)
		} else {
			b := []byte("server:" + msgStr)
			nsConn.Conn.Socket().WriteText(b, 0)
		}
		return nil
	},
}

var connect = func(c *websocket.Conn) error {
	log.Printf("[%s] client Connected to server!", c.ID())
	ctx := websocket.GetContext(c)
	log.Printf("[%s] client url: %s", c.ID(), ctx.FullRequestURI())
	c.Socket().WriteText([]byte("-----------connect"), 0)

	return nil
}

var disconnect = func(c *websocket.Conn) {
	log.Printf("[%s] Disconnected client from server", c.ID())
	c.Socket().WriteText([]byte("-----------disconnect"), 0)
}

func custom(ctx iris.Context) {
	if ctx.IsStopped() {
		return
	}
	page := ctx.URLParamInt64Default("page", 1)
	limit := ctx.URLParamInt64Default("count", 10)
	fmt.Println("-----------------", page, limit)
	ws := websocket.New(gobwas.DefaultUpgrader, RawEvent)
	ws.OnConnect = connect
	ws.OnDisconnect = disconnect
	websocket.Upgrade(ctx, websocket.DefaultIDGenerator, ws)
}

func main() {
	app := iris.New()
	// ws := websocket.New(websocket.DefaultGorillaUpgrader, websocket.Events{
	// 	websocket.OnNativeMessage: func(nsConn *websocket.NSConn, msg websocket.Message) error {

	// 		// return DoSomething(nsConn, msg)
	// 		id := nsConn.Conn.ID()
	// 		log.Printf("Server got: %s from [%s]", msg.Body, id)

	// 		msgStr := string(msg.Body)
	// 		if len(msgStr) > 36 {
	// 			li := bytes.Split(msg.Body, []byte(":"))
	// 			log.Printf("client got id: %s", string(li[0]))
	// 			log.Printf("client got msg: %s", string(li[1]))
	// 			li2 := [][]byte{li[0], []byte("server"), li[1]}
	// 			b := bytes.Join(li2, []byte(":"))
	// 			nsConn.Conn.Socket().WriteText(b, 0)
	// 		} else {
	// 			b := []byte("server:" + msgStr)
	// 			nsConn.Conn.Socket().WriteText(b, 0)
	// 		}
	// 		return nil
	// 	},
	// })

	// ws.OnConnect = func(c *websocket.Conn) error {
	// 	log.Printf("[%s] client Connected to server!", c.ID())
	// 	return nil
	// }

	// ws.OnDisconnect = func(c *websocket.Conn) {
	// 	log.Printf("[%s] Disconnected client from server", c.ID())
	// }
	app.Get("/ws", custom)
	app.Listen(":8888")
}
