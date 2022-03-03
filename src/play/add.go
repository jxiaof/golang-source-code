/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2021-12-30 16:25:06
 * @LastEditTime: 2022-02-24 10:23:46
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	li := make([]int, 0)
	li = append(li, 1)
	li = append(li, 2)
	fmt.Println(li)
	li2 := []int{9999, 0, 0, 0}
	copy(li[1:], li2)
	fmt.Println(li)
	fmt.Println(li2)

	test()
}

func test() {
	defer func() {
		fmt.Println("---------- defer func ------------")
	}()
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Second * 2)
		if i == 5 {
			panic("error")
		}
	}
}
