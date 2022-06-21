/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-03-08 14:41:15
 * @LastEditTime: 2022-03-08 17:36:44
 */

package main

import "fmt"

type a struct {
	name string
	age  int
}

func main() {
	fmt.Println("----")
	var obj a
	test(&obj)
	test(nil)
	fmt.Println("----")
	li := make([]int, 3)
	testc(li)
	fmt.Println(li)
	li2 := [3]int{1, 2, 3}
	testd(li2)
	fmt.Println(li2)
}

func test(obj interface{}) {
	fmt.Println(obj)
	testb(obj)
}

func testb(prt interface{}) {
	fmt.Println(prt)
}

func testc(li []int) {
	li[2] = 100
	fmt.Println(li)
}

func testd(li [3]int) {
	li[2] = 100
	fmt.Println(li)
}
