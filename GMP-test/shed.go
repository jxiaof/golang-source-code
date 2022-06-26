/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-14 01:25:33
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-26 15:14:52
 */
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	runtime.Gosched()
	fmt.Println(s)
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(s)
	// }
}
func main() {
	go say("world")
	// runtime.Gosched()用于让出CPU时间片。
	// runtime.Gosched()
	say("hello")
	// time.Sleep(time.Second)
}
