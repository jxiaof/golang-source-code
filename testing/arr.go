/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-20 10:16:00
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-20 23:37:21
 */
package main

import "fmt"

func main() {

	a := [...]int{1, 2, 3} // ... 会自动计算数组长度
	b := a
	a[0] = 100
	fmt.Println(a, b) // [100 2 3] [1 2 3]

	c := [...]int{1, 2, 3}
	d := &c
	(*d)[0] = 100
	fmt.Println(c, *d) // [100 2 3] [100 2 3]
}
