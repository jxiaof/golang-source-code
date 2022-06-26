/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-27 00:30:47
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-27 00:31:03
 */
package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
	"time"
)

func main() {
	f, err := os.Create("trace.out")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
	}
	err = trace.Start(f)
	defer trace.Stop()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello world !")
	}
	fmt.Println("cpus:", runtime.NumCPU())
}
