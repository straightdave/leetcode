package main

import (
	"fmt"
)

func main() {
	a := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(a))
}

// solution 0: brutal (it works in leetcode!)
func maxArea(height []int) int {
	max := 0

	for i := 0; i < len(height)-1; i++ {
		for j := i + 1; j < len(height); j++ {
			w := j - i
			h := min(height[i], height[j])
			area := w * h
			if area > max {
				max = area
			}
		}
	}

	return max
}

// solution 1: TBD
func maxArea1(height []int) int {

	return 0
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
