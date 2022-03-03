package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	fmt.Println(head.Next)
	r := make([]int, 1)
	for head != nil {
		head = head.Next
		if len(r) == 0 {
			r = append(r, head.Val)
		} else {
			copy(r[:], r[1:])
			r[0] = 9
		}
	}
	return r
}
