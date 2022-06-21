/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-06 17:25:38
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-06 17:28:08
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful/v3"
)

func main() {
	webService := new(restful.WebService)

	webService.Route(webService.GET("/hello").To(hello))

	restful.Add(webService)
	log.Print("Start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	resp.Header().Set("X-Content-Type-Options", "nosniff")
	_, _ = fmt.Fprintf(resp, "world1\n")
	resp.Flush()
	time.Sleep(1 * time.Second)
	_, _ = fmt.Fprintf(resp, "world2\n")
	resp.Flush()
	time.Sleep(1 * time.Second)
	_, _ = fmt.Fprintf(resp, "world3\n")
	resp.Flush()
	time.Sleep(1 * time.Second)
	_, _ = fmt.Fprintf(resp, "world4\n")
}
