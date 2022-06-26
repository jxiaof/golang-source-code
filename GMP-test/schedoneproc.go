/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-26 20:35:51
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-27 01:15:03
 */
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	for i := 1; i < 259; i++ {
		go func(i int) {
			fmt.Println("----->", i)
		}(i)
	}

	time.Sleep(time.Millisecond * 100)

}

// 258 = 256 + 1 + 1
// 257 + localQueue / 2 (129, 130, 131 ... 256 )     , globalQueue  (1, 2, 3 ... 128, 257 )
// schedtick = 61
