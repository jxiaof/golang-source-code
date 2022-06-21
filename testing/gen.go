/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-19 15:37:28
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-19 15:37:31
 */

package main

import (
	"math/rand"
	"time"
)

func generateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
