package main

import (
	"fmt"
)

func main() {
	a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(a))
}

func maxSubArray(nums []int) int {
	const MaxUint = ^uint(0)
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1

	max := MinInt
	loc := 0

	for _, n := range nums {
		loc += n

		if loc > max {
			max = loc
		}

		if loc < 0 {
			loc = 0
		}
	}

	return max
}
