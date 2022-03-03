package main

import (
	"fmt"
	"sync"
)

func main() {
	var num int = 3
	dog := make(chan struct{})
	cat := make(chan struct{})
	fox := make(chan struct{})
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go func(n int) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-dog
			fmt.Println("dog")
			cat <- struct{}{}
		}
	}(num)
	go func(n int) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			<-cat
			fmt.Println("cat")
			fox <- struct{}{}
		}
	}(num)
	go func(n int) {
		defer wg.Done()
		for i := 0; i < n; i++ {
			fmt.Println("------", i)
			<-fox
			fmt.Println("fox")
			if i < n-1 {
				dog <- struct{}{}
			}
		}
	}(num)
	dog <- struct{}{}
	wg.Wait()
	fmt.Println("finished!!!")
}
