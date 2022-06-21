/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-10 10:32:35
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-10 10:45:40
 */
package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup
	defer func() {
		wg.Wait()
	}()

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		i := i
		go func() {
			defer wg.Done()
			worker(i)
		}()
	}

}
