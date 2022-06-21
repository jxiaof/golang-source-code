/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-10 09:56:33
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-24 16:38:27
 */
package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func goRoutineD(ch chan int, i int) {
	time.Sleep(time.Second * 3)
	ch <- i
}
func goRoutineE(chs chan string, i string) {
	time.Sleep(time.Second * 3)
	chs <- i
}

func foo() <-chan int {
	ch := make(chan int)
	defer close(ch)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()

	return ch
}
func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

type SClusterInfo struct {
	Key         string `json:"_id" bson:"_id"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	GrpcURL     string `json:"grpc_url" bson:"grpc_url"`
	RestURL     string `json:"rest_url" bson:"rest_url"`
	// CallMode    TCluCallMode `json:"call_mode,omitempty" bson:"call_mode,omitempty"`
	CreTime time.Time `json:"cre_time" bson:"cre_time"`
	Token   string    `json:"token,omitempty" bson:"token,omitempty"`
}

var CTX_KEY_CLUSTER = "cluster"

func main() {
	// ch := make(chan int, 5)
	// chs := make(chan string, 5)

	// go goRoutineD(ch, 5)
	// go goRoutineE(chs, "ok")

	// select {
	// case msg := <-ch:
	// 	fmt.Println(" received the data ", msg)
	// case msgs := <-chs:
	// 	fmt.Println(" received the data ", msgs)
	// 	// default:
	// 	//     fmt.Println("no data received ")
	// 	//     time.Sleep(time.Second * 1)
	// }

	// ch2 := foo()
	// for v := range ch2 {
	// 	fmt.Println(v)
	// }
	// time.Sleep(time.Second * 3)

	// naturals := make(chan int)
	// squares := make(chan int)
	// go counter(naturals)
	// go squarer(squares, naturals)
	// printer(squares)

	ch := make(chan int, 1)

	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 1)
			ch <- i
		}
		close(ch)
	}()
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("start")
		t := time.NewTicker(time.Second * 3)
		defer t.Stop()
		li := make([]int, 0)
		var a string
		p := func() {
			fmt.Println(li)
			a = ""
			li = make([]int, 0)
		}
		for {
			// go func() {
			// 	fmt.Println("ticker")
			// 	<-t.C
			// 	p()
			// }()
			select {
			case <-t.C:
				p()

			default:
				fmt.Println("default")
				time.Sleep(time.Millisecond * 100)
			}

			select {
			case val, ok := <-ch:
				if ok {
					// fmt.Println(val)
					li = append(li, val)
					a += fmt.Sprintf("get val: %d", val)
				} else {
					fmt.Println("channel closed")
					p()
					return
				}
			case <-time.After(time.Second * 11):
				fmt.Println("timeout")
				return
			default:
				fmt.Println("no data received")
				time.Sleep(time.Millisecond * 500)
			}
			fmt.Println("goroutine number:", runtime.NumGoroutine())

		}
	}()

	ctx := context.WithValue(context.TODO(), CTX_KEY_CLUSTER, &SClusterInfo{Key: "zhai-dev"})
	cluster, ok := ctx.Value(CTX_KEY_CLUSTER).(*SClusterInfo)
	if !ok || cluster == nil || cluster.Key == "" {
		fmt.Println("get cluster info failed")
	}
	fmt.Println(cluster.Key)
	// time.Sleep(time.Second * 12)
	wg.Wait()
	fmt.Println("------goroutine number:", runtime.NumGoroutine())
}
