/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-19 14:18:08
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-19 14:18:12
 */
// fib.go
package main

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}
