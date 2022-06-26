/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-26 23:49:14
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-27 01:05:54
 */
package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
)

func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	runtime.GOMAXPROCS(8)
	wg := &sync.WaitGroup{}
	go func() {
		wg.Add(1)
		fmt.Println("Hello")
		wg.Done()
	}()
	go func() {
		wg.Add(1)
		fmt.Println("World")
		wg.Done()
	}()
	wg.Wait()
}
