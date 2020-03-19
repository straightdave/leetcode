package main

import (
	"fmt"
)

func main() {
	a := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares(a))
}

func sortedSquares(A []int) []int {
	var s []int

	for _, n := range A {
		s = append(s, n*n)
	}

	return qsort(s)
}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	var (
		smaller []int
		greater []int
		pivot   = a[0]
	)

	for _, n := range a[1:] {
		if n < pivot {
			smaller = append(smaller, n)
		} else {
			greater = append(greater, n)
		}
	}

	smaller = qsort(smaller)
	greater = qsort(greater)
	return append(append(smaller, pivot), greater...)
}
