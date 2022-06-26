package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/kataras/iris/v12"
)

type clientPage struct {
	Title string
	Host  string
}

const (
	// Time allowed to write the file to the client.
	writeWait = 10 * time.Second
	// Time allowed to read the next pong message from the client.
	pongWait   = 60 * time.Second
	pingPeriod = 5 * time.Second
	msgPeriod  = 1 * time.Second
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
}

func WsSvr(ctx iris.Context) {
	fmt.Println("WS SVR ------> :", ctx.Request().RemoteAddr)
	ws, err := upgrader.Upgrade(ctx.ResponseWriter(), ctx.Request(), nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}
		return
	}

	go Writer(ws)
	Reader(ws)
}

func Writer(ws *websocket.Conn) {
	ping := time.NewTicker(pingPeriod)
	msg := time.NewTicker(msgPeriod)
	defer func() {
		ping.Stop()
		msg.Stop()
		ws.Close()
	}()
	for {
		select {
		case <-msg.C:
			timeStr := time.Now().Format("2006-01-02 15:04:05")
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.TextMessage, []byte(timeStr)); err != nil {
				return
			}
		case <-ping.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func Reader(ws *websocket.Conn) {
	defer ws.Close()
	ws.SetReadLimit(2048)
	ws.SetReadDeadline(time.Now().Add(pongWait))
	ws.SetPongHandler(func(string) error { ws.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, _, err := ws.ReadMessage()
		if err != nil {
			break
		}
	}
}
func Hello(ctx iris.Context) {
	ctx.JSON("hello")
}

func main() {
	app := iris.New()
	// app.RegisterView(iris.HTML("./views", ".html"))
	// app.HandleDir("/js", iris.Dir("./static/js"))
	// app.Get("/", func(ctx iris.Context) {
	// 	ctx.View("index.html", clientPage{"---- -- -- - Client Page - - -- - -------", "localhost:9999"})
	// })
	// use logger
	app.Any("/ws", WsSvr)
	app.Any("/", Hello)
	fmt.Println(app.Router)
	app.Listen(":9999")
}
