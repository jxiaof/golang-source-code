/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-20 09:36:40
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-20 10:06:11
 */
package main

import (
	"fmt"
	"math/rand"
	"os"
	"runtime/pprof"
	"time"

	ppf "github.com/pkg/profile"
)

func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

func GenerateWithCap(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}

func main() {
	f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	n := 10
	for i := 0; i < 5; i++ {
		nums := GenerateWithCap(n)
		bubbleSort(nums)
		n *= 10
	}

	fmt.Println("----------------------------------------------------")

	defer ppf.Start().Stop()
	nums := GenerateWithCap(1000)
	bubbleSort(nums)
}
