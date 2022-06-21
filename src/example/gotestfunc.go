/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-11 14:55:13
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-11 15:00:15
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	ch := foo()
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
}

func foo() <-chan string {
	fmt.Println("foo start")
	ch := make(chan string)
	go func() {
		defer close(ch)
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			ch <- "bar start"
		}
	}()
	fmt.Println("foo end")
	return ch
}
