package main

import "fmt"

func main() {
	r := foo(4)
	fmt.Println(r)
}

func foo(n int) int {
	// 输入一个正整数,输出正整数平方根
	var y int = 0
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}

	if y*y < n {
		y++
	} else {
		return y
	}
}

func bar(li []int) int {
	// 输入一个数组,输出数组中两次出现的值
	var a int = 0
	var b int = 0
	for i := 0; i < len(li); i++ {
		if li[i] == a {
			b = a
			a = li[i]
		} else if li[i] == b {
			b = li[i]
		}
	}
	return b
}
