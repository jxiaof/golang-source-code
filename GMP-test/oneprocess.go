/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-20 23:37:56
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-26 23:49:00
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)

	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("----->", i)
		}(i)
	}
	wg.Wait()
}
