/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-04-20 16:42:54
 * @LastEditTime: 2022-04-25 17:21:49
 */

package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
	str := "hello;world"
	fmt.Println(len(strings.Split(str, ";")))
}
⁄