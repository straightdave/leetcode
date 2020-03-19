package main

import (
	"fmt"
)

func main() {
	// a := [][]int{
	// 	[]int{1, 1},
	// 	[]int{3, 4},
	// 	[]int{-1, 0},
	// }
	a := [][]int{
		[]int{3, 2},
		[]int{-2, 2},
	}

	fmt.Println(minTimeToVisitAllPoints(a))
}

func minTimeToVisitAllPoints(points [][]int) int {
	sum := 0

	if len(points) == 0 {
		return 0
	}

	lastPoint := points[0]
	for _, point := range points[1:] {
		sum += min(lastPoint, point)
		lastPoint = point
	}

	return sum
}

func min(a, b []int) int {
	horSub := abs(a[0] - b[0])
	verSub := abs(a[1] - b[1])

	if horSub > verSub {
		return horSub
	} else {
		return verSub
	}
}

func abs(n int) int {
	if n >= 0 {
		return n
	}
	return -n
}
