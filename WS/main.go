/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-21 11:13:20
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-22 10:20:27
 */
package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// run starts a http.Server for the passed in address
// with all requests handled by echoServer.
func run() error {

	addr := "localhost:10000"

	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	log.Printf("listening on http://%s", addr)

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

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	return s.Shutdown(ctx)
}

func main() {
	log.SetFlags(0)

	err := run()
	if err != nil {
		log.Fatal(err)
	}
}
