/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-24 15:52:49
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-24 15:54:16
 */
package main

import "fmt"

// switch test
func main() {
	// switch test
	var a interface{}
	a = 1
	switch a.(type) {
	case int:
		fmt.Println("a is int")
	case string:
		fmt.Println("a is string")
	default:
		fmt.Println("a is other")
	}

	// // switch test
	// switch {
	// case a.(type) == int:
	// 	fmt.Println("a is int")
	// case a.(type) == string:
	// 	fmt.Println("a is string")
	// default:
	// 	fmt.Println("a is other")
	// }
}
