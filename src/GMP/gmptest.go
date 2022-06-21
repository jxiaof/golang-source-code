/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-12 23:20:55
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-12 23:22:47
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
		fmt.Println("GMP")
	}
	fmt.Println("cpus:", runtime.NumCPU())
}
