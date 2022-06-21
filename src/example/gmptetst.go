/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-10 10:19:28
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-10 10:27:10
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go func() {
		fmt.Println("hello world")
		time.Sleep(time.Second * 5)
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("start:", runtime.NumGoroutine())
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("hello world 22222")
			time.Sleep(time.Second * 1)
		}
	}()
	time.Sleep(time.Second * 1)
	fmt.Println("end:", runtime.NumGoroutine())
	// select {}
}
