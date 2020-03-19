package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	r := &TreeNode{
		Val: 10,
		Left: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
		Right: &TreeNode{
			Val: 15,
			Right: &TreeNode{
				Val: 18,
			},
		},
	}

	fmt.Println(rangeSumBST(r, 7, 15))
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	var (
		sum = 0
		f   func(*TreeNode)
	)

	f = func(n *TreeNode) {
		if n == nil {
			return
		}

		if n.Val >= L && n.Val <= R {
			sum += n.Val
		}

		f(n.Left)
		f(n.Right)
	}

	f(root)
	return sum
}
