/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-01-21 11:30:46
 * @LastEditTime: 2022-01-21 11:30:49
 */
package main

import (
	"fmt"

	"github.com/golang-module/carbon"
)

func main() {
	fmt.Println("Hello, 世界")
	s := fmt.Sprintf("%s", carbon.Now())  // 2020-08-05 13:14:15
	s2 := carbon.Now().ToDateTimeString() // 2020-08-05 13:14:15
	fmt.Println(s, s2)
}
