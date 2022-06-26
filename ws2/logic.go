/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-24 10:32:15
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-24 10:45:07
 */
package main

import (
	"bytes"
	"log"

	"github.com/kataras/iris/v12/websocket"
)

func DoSomething(nsConn *websocket.NSConn, msg websocket.Message) error {
	// for i := 0; i < 100; i++ {

	// 	time.Sleep(time.Second*5)

	// }
	// server recv msg
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
}
