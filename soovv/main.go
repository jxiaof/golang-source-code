/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-04-11 09:44:22
 * @LastEditTime: 2022-04-11 10:24:25
 */

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("Hello World!")
	ch1 := make(chan string)
	ch2 := make(chan string)
	go run(ch1, ch2)
	for {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
			if msg2 == "Finished" {
				print("Finished os.Exit(0)\n")
				os.Exit(0)
			}
		case <-time.After(3 * time.Second):
			fmt.Println("Timeout! os.Exit(1)\n")
			os.Exit(1)
			// default:
		}
	}
}

func run(ch1, ch2 chan string) {
	for i := 0; i < 12; i++ {
		if i < 3 {
			ch2 <- "Padding!"
		}
		if i == 5 {
			ch2 <- "Running"
		}
		if i == 12 {
			// time.Sleep(time.Second * 5)
			ch2 <- "Finished"
		}
		if i == 10 {
			fmt.Println("break")
			return
		}
		ch1 <- "Hello World!"
		time.Sleep(time.Second)
	}
}
