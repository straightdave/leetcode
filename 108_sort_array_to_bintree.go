//
// Not Finished Yet
//

package main

import (
	"fmt"
)

func main() {

	a := []int{-10, -3, 0, 5, 9}
	r := sortedArrayToBST()

}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) < 1 {
		return nil
	}

	if len(nums) == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}

	if len(nums) == 2 {
		return &TreeNode{
			Val: nums[1],
			Left: &TreeNode{
				Val: nums[0],
			},
		}
	}

	rootIndex := len(nums) / 2
	lroot := sortedArrayToBST(nums[0:rootIndex])
	rroot := sortedArrayToBST(nums[rootIndex+1:])

	return &TreeNode{
		Val:   nums[rootIndex],
		Left:  lroot,
		Right: rroot,
	}
}
