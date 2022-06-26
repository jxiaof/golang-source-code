/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-23 16:39:27
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-24 16:53:07
 */
package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kataras/iris/v12/websocket"
	"github.com/kataras/neffos/gobwas"
)

var nativeEvent = websocket.Events{
	websocket.OnNativeMessage: func(c *websocket.NSConn, msg websocket.Message) error {
		log.Printf("client got: %s from [%v]", msg.Body, c.Conn.ID())
		// fmt.Printf("client got: %s from [%s]", msg.Body, c.Conn.ID())
		li := bytes.Split(msg.Body, []byte(":"))
		log.Printf("client got id: %s", string(li[0]))
		log.Printf("client got msg: %s", string(li[1]))

		// nsConn.Conn.Server().Broadcast(nsConn, msg)
		// m := "echo :" + string(msg.Body)
		// mg := websocket.Message{
		// 	Body:     []byte(m),
		// 	IsNative: true,
		// }/webhp

		// nsConn.Conn.Write(mg)
		return nil
	},
}

func main() {
	ctx := context.Background()
	endpoint := "ws://localhost:8888/ws"

	// username := "uuid"
	// dialer := websocket.GobwasDialer(websocket.GobwasDialerOptions{Header: websocket.GobwasHeader{"X-Username": []string{username}}})
	dialer := gobwas.DefaultDialer
	client, err := websocket.Dial(ctx, dialer, endpoint, nativeEvent)
	if err != nil {
		panic(err)
	}
	defer client.Close()

	c, err := client.Connect(ctx, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect id", c.Conn.ID(), "----", c.String(), c.Conn == nil)
	c.Conn.Socket().WriteText([]byte("ping"), 0)
	time.Sleep(time.Second * 5)
	c.Conn.Socket().WriteText([]byte("pong"), 0)
	time.Sleep(time.Second * 5)
	mg := websocket.Message{
		Body:     []byte("hello"),
		IsNative: true,
	}
	time.Sleep(time.Second * 5)
	c.Conn.Write(mg)
	for i := 0; i < 100; i++ {
		c.Conn.Socket().WriteText([]byte("ping"), 0)
		time.Sleep(time.Second * 5)
	}
}
