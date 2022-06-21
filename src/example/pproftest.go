/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-24 17:20:33
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-25 14:55:19
 */
package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	// we need a webserver to get the pprof webserver
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	fmt.Println("hello world")
	http.HandleFunc("/hello", myHandler)
	var wg sync.WaitGroup
	wg.Add(1)
	go leakyFunction(wg)
	wg.Wait()
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!\n")
}

func leakyFunction(wg sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "magical pandas")
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}

//  func PprofWeb() {
// 	err := http.ListenAndServe(":9909", nil)
// 	if err != nil {
// 	   panic(err)
// 	}
//  }
