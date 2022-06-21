/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-07 11:53:11
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-07 12:06:07
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	str := foo()
	fmt.Println(str)
	bar()
	time.Sleep(time.Second * 5)
}

func foo() string {
	fmt.Println("foo start")
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			fmt.Println("bar start")
		}
	}()
	fmt.Println("foo end")
	return "hello foo"
}

func bar() {
	time.Sleep(time.Second * 2)
	fmt.Println("bar start")
}
