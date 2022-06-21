/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-19 14:18:08
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-19 15:33:02
 */
// fib.go
package main

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "test1",
			args: args{
				n: 10,
			},
			want: 55,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib(30) // run the Fib function b.N times
	}
}
