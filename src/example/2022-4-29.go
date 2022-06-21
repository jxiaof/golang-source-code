/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-04-29 17:32:52
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-04-29 17:32:53
 */
// You can edit this code!
// Click here and start typing.
package main

import "fmt"

// THardwarePlat 硬件平台
type THardwarePlat string

const (
	// 硬件平台
	THwUnknown THardwarePlat = ""
	THwCPU     THardwarePlat = "CPU"
	THwGPU     THardwarePlat = "GPU"
)

func main() {

	fmt.Println("Hello, 世界", THardwarePlat("CPU"), THardwarePlat("CPU") == THwCPU)
}
