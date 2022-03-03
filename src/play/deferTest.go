/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-02-28 10:00:45
 * @LastEditTime: 2022-02-28 10:45:28
 */

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	t1 := time.Now()
	str := "hello"
	defer fmt.Println(str)
	// defer fmt.Println(str)

	fmt.Println(str)
	str = "world"
	defer fmt.Println(str)
	time.Sleep(time.Second * 3)
	fmt.Println(time.Since(t1).Seconds(), time.Since(t1) > time.Second*2, time.Since(t1) > 2*1)

	val, err := strconv.Atoi("1800")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(val, val+1)
}
