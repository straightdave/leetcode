package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	h := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 0,
			Next: &ListNode{
				Val:  1,
				Next: nil,
			},
		},
	}

	fmt.Println(getDecimalValue(h))
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) int {
	var bin []int
	p := head
	for p != nil {
		bin = append(bin, p.Val)
		p = p.Next
	}

	l := len(bin) - 1
	sum := 0
	for i, b := range bin {
		if b == 0 {
			continue
		}
		sum += (1 << (l - i))
	}

	return sum
}
