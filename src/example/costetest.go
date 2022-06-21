/*
 * @Descripttion:
 * @version:
 * @Author: hujianghong
 * @Date: 2022-05-17 16:33:28
 * @LastEditors: hujianghong
 * @LastEditTime: 2022-05-19 11:28:26
 */
package main

import "fmt"

type WorkflowPhase string

const (
	WorkflowUnknown   WorkflowPhase = ""
	WorkflowPending   WorkflowPhase = "Pending" // pending some set-up - rarely used
	WorkflowRunning   WorkflowPhase = "Running" // any node has started; pods might not be running yet, the workflow maybe suspended too
	WorkflowSucceeded WorkflowPhase = "Succeeded"
	WorkflowFailed    WorkflowPhase = "Failed" // it maybe that the the workflow was terminated
	WorkflowError     WorkflowPhase = "Error"
)

func main() {
	// fmt.Println(string(WorkflowRunning))
	var tmp string
	foo := func(p string) {
		if tmp == p {
			fmt.Println("no change, pass")
			return
		}
		tmp = p
		fmt.Println(p)

	}

	foo(string(WorkflowPending))
	foo(string(WorkflowRunning))
	foo(string(WorkflowRunning))
	foo(string(WorkflowRunning))
	foo(string(WorkflowSucceeded))

}
