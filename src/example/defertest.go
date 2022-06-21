/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-12 17:33:02
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-16 20:25:08
 */
package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

// defer is used to execute a function after the surrounding function returns
func main() {
	defer fmt.Println("---- end -----")
	defer fmt.Println("this is the main func end")

	log.Info("------------ start baz ------------")
	var flagBaz bool
	var bazCount int
	baz := baz()
	for !flagBaz {
		select {
		case val, ok := <-baz:
			if !ok {
				flagBaz = true
				log.Info("baz done")
				break
			}
			bazCount++
			log.Infof("baz recv: %d", val)
		case <-time.After(time.Second * 5):
			log.Info("baz timeout")
			flagBaz = true
			break
		}
	}
	log.Infof("baz count: %d", bazCount)

	log.Info("------------ end baz ------------")

	log.Info("------------ start foo ------------")
	var flag bool
	f := foo()
	for !flag {
		select {
		case i, ok := <-f:
			if !ok {
				log.Info("channel Closed!")
				flag = true
				break
			}
			log.Infof("recv: %d", i)
		case <-time.After(time.Second * 5):
			log.Info("timeout")
			flag = true
			break
		}
	}
	log.Info("------------ end foo ------------")

	log.Info("------------ start bar ------------")
	b := bar()
	for {
		select {
		case i, ok := <-b:
			if !ok {
				log.Info("channel Closed!")
				return
			}
			log.Infof("recv: %d", i)
			// time.Sleep(time.Second)
		case <-time.After(time.Second * 5):
			log.Info("timeout")
			return
		}
	}
	log.Info("------------ end bar ------------")

}

func foo() <-chan int {
	loopNum := 3
	ch := make(chan int, 1)
	done := make(chan struct{})
	wait := func(ch chan int, done chan struct{}, i int) {
		j := 0
		defer func() {
			log.Info("close channel")
			close(done) // 接收端关闭,不是一个好的做法. 可以设置多个channel,一对一关闭
			time.Sleep(time.Second * time.Duration(loopNum))
			close(ch)
		}()
		for {
			select {
			case <-done:
				log.Infof("wait: %d", i)
				j++
				if j == i {
					return
				}
			case <-time.After(time.Second * 5):
				log.Info("timeout")
				return

			}
		}
	}

	send := func() {
		for i := 0; i < loopNum; i++ {
			ch <- i
			log.Infof("send: %d", i)
			time.Sleep(time.Second)
		}
		done <- struct{}{}
	}

	for j := 0; j < loopNum; j++ {
		go send()
	}
	go wait(ch, done, loopNum)

	return ch
}

func bar() <-chan int {
	loopNum := 3
	ch := make(chan int, 1)
	wg := &sync.WaitGroup{}
	lock := sync.Mutex{}

	send := func() {
		lock.Lock()
		defer lock.Unlock()
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			for i := 0; i < num; i++ {
				ch <- i
				time.Sleep(time.Second)
				log.Infof("send: %d", i)
			}
		}(loopNum)

	}

	wait := func(ch chan int, wg *sync.WaitGroup) {
		log.Info("wait")
		wg.Wait()
		time.Sleep(time.Second * time.Duration(loopNum))
		log.Info("close channel")
		close(ch)
	}

	for j := 0; j < loopNum; j++ {
		// wg.Add(1)
		go send()
	}

	go wait(ch, wg)

	return ch
}

func baz() <-chan int {
	// res:
	// 	for i := 0; i < 10; i++ {
	// 		if i == 5 {
	// 			break res
	// 		}
	// 		log.Info("baz: ", i)
	// 	}
	// 	log.Info("baz done")

	log.Info("baz start gorutine num: ", runtime.NumGoroutine())

	ch := make(chan int, 1)
	wg := &sync.WaitGroup{}
	lock := sync.Mutex{}
	loopNum := 3

	send := func(num int) {
		lock.Lock()
		defer lock.Unlock()
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			defer log.Infof("send done, num: %d", num)
			for i := 1; i <= num; i++ {
				ch <- 1
				time.Sleep(time.Second)
				log.Infof("send: %d", i)
			}
		}(num)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= loopNum; i++ {
			log.Info("send baz: ", i)
			send(i)
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := 0
			for i := 1; i <= loopNum; i++ {
				log.Info("send baz1: ", i)
				send(i)
				tmp += i
			}
			log.Info("send baz1 done: ", tmp)
		}()

		log.Info("another goroutine")
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 1; i <= loopNum; i++ {
				log.Info("send baz2: ", i)
				send(i)
			}
		}()
		log.Info("all goroutine running")
	}()

	go func() {
		log.Info("wait")
		wg.Wait()
		time.Sleep(time.Second * time.Duration(loopNum))
		log.Info("close channel")
		close(ch)
		log.Info("baz end gorutine num: ", runtime.NumGoroutine())
	}()

	// go wait()
	log.Info("baz running gorutine num: ", runtime.NumGoroutine())
	return ch
}
