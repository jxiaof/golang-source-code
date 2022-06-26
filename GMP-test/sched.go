/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-06-14 01:25:33
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-06-26 23:00:20
 */
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	runtime.Gosched()
	fmt.Println(s)
}
func main() {
	go say("world")
	// runtime.Gosched()
	say("hello")
}

// runtime.main() -> main.main -> newprco(0, say) (run_next)-> exit -> goexit() 进程结束

// (sysmon 监控) time.Sleep -> P.timer (G_waiting) -> G_running -> localQueue

//  channel -> runtime.gopark -> G_running -> lock -> G_waiting -> G_running -> localQueue -> cur.G
