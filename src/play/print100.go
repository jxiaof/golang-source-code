package main

import (
	"fmt"
	"runtime"
	"sync"
)

const (
	NUM = 5
)

func main() {
	test1()
	// test2()
	// test3()
	// test4()
	// test5()
}

func test1() {
	wg := &sync.WaitGroup{}
	ch := make(chan struct{})
	wg.Add(2)
	go print1(ch, wg)
	go print2(ch, wg)
	wg.Wait()
}

func test2() {
	ch := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer func() {
			wg.Done()
			close(ch)
		}()
		for i := 1; i < NUM; i += 2 {
			//奇数
			ch <- struct{}{}
			fmt.Println("线程1打印:", i)
			<-ch

		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i < NUM; i += 2 {
			//偶数
			fmt.Println("线程2打印:", i)
			<-ch
			ch <- struct{}{}

		}
	}()
	wg.Wait()
}

func test3() {
	var ch = make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= NUM; i += 2 {
			fmt.Println("-.--.------", i)
			ch <- struct{}{} //不能与上一行交换位置
			<-ch
		}
	}()
	go func() {
		defer wg.Done()
		for i := 2; i <= NUM; i += 2 {
			<-ch
			fmt.Println("+++.+++.+++", i)
			ch <- struct{}{}
		}
	}()
	wg.Wait()
}
func print1(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 2; i <= NUM; i += 2 {
		fmt.Println("-------", i)
		ch <- struct{}{}
		<-ch
	}

}

func print2(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= NUM; i += 2 {
		fmt.Println("+++++++", i)
		<-ch
		ch <- struct{}{}
	}
}

func test4() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= NUM; i++ {
			ch <- i
			n := <-ch
			fmt.Println("-------", 2*n)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 1; i <= NUM; i++ {
			n := <-ch
			ch <- n
			fmt.Println("+++++++", 2*n-1)
		}
	}()
	wg.Wait()
}

func test5() {
	//设置可同时使用的CPU核数为1
	var wg sync.WaitGroup
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= NUM; i += 2 {
			//奇数
			fmt.Println("线程1打印:", i)
			//让出cpu
			runtime.Gosched()
		}
	}()
	go func() {
		defer wg.Done()
		runtime.Gosched()
		for i := 2; i <= NUM; i += 2 {
			//偶数
			fmt.Println("线程2打印:", i)
			//让出cpu
			runtime.Gosched()
		}
	}()
	wg.Wait()
}
