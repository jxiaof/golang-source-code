/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-12 09:56:11
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-12 09:59:56
 */
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	// 测试context
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("程序结束")
				return
			case <-time.After(time.Second * 1):
				fmt.Println("1秒后执行")
			}
		}
	}()
	time.Sleep(time.Second * 2)
	fmt.Println("取消任务")
	cancel()
	fmt.Println("-------------------------")
	time.Sleep(time.Second * 3)
}
