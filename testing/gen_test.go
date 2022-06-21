/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-19 15:37:28
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-19 15:58:59
 */

package main

import (
	"testing"
)

// func Test_generateWithCap(t *testing.T) {
// 	type args struct {
// 		n int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []int
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "genCap_test",
// 			args: args{
// 				n: 10,
// 			},
// 			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := generateWithCap(tt.args.n); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("generateWithCap() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func BenchmarkGenerateWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateWithCap(10) // run the Fib function b.N times
	}
}

func BenchmarkGenerateWithCap1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateWithCap(1000) // run the Fib function b.N times
	}
}

func BenchmarkGenerateWithCap1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateWithCap(1000000) // run the Fib function b.N times
	}
}

// func Test_generate(t *testing.T) {
// 	type args struct {
// 		n int
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []int
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "gen_test",
// 			args: args{
// 				n: 10,
// 			},
// 			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := generate(tt.args.n); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("generate() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(10) // run the Fib function b.N times
	}
}

func BenchmarkGenerate1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(1000) // run the Fib function b.N times
	}
}

func BenchmarkGenerate1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(1000000) // run the Fib function b.N times
	}
}
