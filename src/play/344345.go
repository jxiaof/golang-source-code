/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-03-03 20:57:41
 * @LastEditTime: 2022-03-03 21:02:57
 */

// 两个channel，交替打印
package main

import (
	"fmt"
	"time"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	done := make(chan bool)

	go func(a chan int) {
		for i := 0; i < 5; i += 2 {
			a <- i
			fmt.Println("a", i)
		}
		done <- true
	}(a)

	go func(b chan int) {
		for i := 1; i < 5; i += 2 {
			b <- i
			fmt.Println("a", i)
		}
		done <- true
	}(b)

	time.Sleep(time.Second)
}
