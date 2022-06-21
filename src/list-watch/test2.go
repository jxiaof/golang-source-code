/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-06 17:30:45
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-06 17:31:39
 */
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

var myChan = make(chan string)

func main() {
	webService := new(restful.WebService)

	webService.Route(webService.GET("/hello").To(hello))
	webService.Route(webService.GET("/ok/{id}").To(ok))

	restful.Add(webService)
	log.Print("Start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	resp.Header().Set("X-Content-Type-Options", "nosniff")
	_, _ = fmt.Fprintf(resp, "wait:\n")
	resp.Flush()
	for i := range myChan {
		fmt.Println(i)
		_, _ = fmt.Fprintf(resp, "id: %s\n", i)
		resp.Flush()
	}
}

func ok(req *restful.Request, resp *restful.Response) {
	myChan <- req.PathParameter("id")
	_, _ = fmt.Fprintf(resp, "OK")
}
