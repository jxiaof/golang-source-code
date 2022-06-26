/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-22 10:42:02
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-22 21:50:05
 */
package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"golang.org/x/time/rate"

	"nhooyr.io/websocket"
)

func Hello() {
	fmt.Println("--------->")
}

type echoServer struct {
	// logf controls where logs are sent.
	logf func(f string, v ...interface{})
}

func (s echoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		fmt.Println("err start server")
		return
	}
	defer c.Close(websocket.StatusInternalError, "the sky is falling")
	l := rate.NewLimiter(rate.Every(time.Millisecond*100), 10)
	for {
		fmt.Println("---------->")
		err = echo(r.Context(), c, l)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
		if err != nil {
			fmt.Println("failed to echo with:", r.RemoteAddr, " err:", err)
			return
		}
	}
	fmt.Println("----------> done")
}

func echo(ctx context.Context, c *websocket.Conn, l *rate.Limiter) error {
	// ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	// defer cancel()

	err := l.Wait(ctx)
	if err != nil {
		return err
	}

	typ, r, err := c.Reader(ctx)
	if err != nil {
		return err
	}

	w, err := c.Writer(ctx, typ)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("failed to io.Copy: %w", err)
	}
	err = w.Close()
	return err
}

func Run() error {
	l, err := net.Listen("tcp", ":9900")
	if err != nil {
		return err
	}
	log.Printf("listening on http://%s", "localhost:9900")

	s := &http.Server{
		Handler: echoServer{
			logf: log.Printf,
		},
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 10,
	}

	errc := make(chan error, 1)
	go func() {
		errc <- s.Serve(l)
	}()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	select {
	case err := <-errc:
		log.Printf("failed to serve: %v", err)
	case sig := <-sigs:
		log.Printf("terminating: %v", sig)

	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()

	return s.Shutdown(ctx)

}

func main() {
	log.SetFlags(0)

	err := Run()
	if err != nil {
		log.Fatal(err)
	}
}
