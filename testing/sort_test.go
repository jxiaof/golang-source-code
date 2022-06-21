/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-20 09:36:40
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-20 09:39:14
 */
package main

import (
	"testing"
)

// func Test_bubbleSort(t *testing.T) {
// 	type args struct {
// 		nums []int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			bubbleSort(tt.args.nums)
// 		})
// 	}
// }

func BenchmarkBubbleSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		nums := generateWithCap(10000)
		b.StartTimer()
		bubbleSort(nums)
	}
}
