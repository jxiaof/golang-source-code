/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-27 00:44:46
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-27 00:47:15
 */
package main

import (
	"net/http"
	"os"
	"runtime/trace"
	"sync"
)

var url = "https://www.google.com"

func main() {

	f, _ := os.Create("traceHttp.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		_, _ = http.Get(url)
		wg.Done()
	}
	wg.Wait()
}
