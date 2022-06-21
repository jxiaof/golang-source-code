/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-20 23:37:56
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-21 00:30:41
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	starTime := time.Now()
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			time.Sleep(time.Second)
			if i == 3 {
				time.Sleep(time.Second * 10)
			}
			fmt.Println("----->", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("time:", time.Since(starTime))
}
