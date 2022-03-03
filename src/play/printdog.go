package main

import (
	"fmt"
	"sync"
)

const (
	LOOP_NUM int = 3
)

func foo(wg *sync.WaitGroup, cha, chb chan struct{}, msg string, end bool) {
	defer wg.Done()
	i := 0
	num := LOOP_NUM
	if end {
		num = LOOP_NUM - 1
	}
	for _ = range cha {
		i += 1
		fmt.Println(msg)
		if i <= num {
			chb <- struct{}{}
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	wg.Add(3)
	a, b, c := make(chan struct{}), make(chan struct{}), make(chan struct{})
	go foo(wg, a, b, "hello", false)
	go foo(wg, b, c, "world", false)
	go foo(wg, c, a, "!", true)
	a <- struct{}{}
	wg.Wait()
	fmt.Println("end")
}
