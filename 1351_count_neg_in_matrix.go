package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		[]int{4, 3, 2, -1},
		[]int{3, 2, 1, -1},
		[]int{1, 1, -1, -2},
		[]int{-1, -1, -2, -3},
	}
	fmt.Println(countNegatives(a))
}

func countNegatives(grid [][]int) int {
	count := 0

	for _, row := range grid {
		for i, n := range row {
			if n < 0 {
				count += len(row) - i
				break
			}
		}
	}

	return count
}
