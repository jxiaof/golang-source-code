/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-25 15:05:44
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-25 16:02:03
 */
package main

import (
	"fmt"

	"github.com/kataras/iris/v12"

	// "github.com/kataras/iris/v12/middleware/pprof"

	_ "runtime/pprof"
)

func main() {
	app := iris.New()
	app.Get("/", func(ctx iris.Context) {
		fmt.Println("hello world")
		ctx.HTML("<h1> Please click <a href='/debug/pprof'>here</a>")
	})
	// app.Any("/debug/pprof/{action:path}", New())
	//                              ___________
	app.Run(iris.Addr(":8080"))

}
