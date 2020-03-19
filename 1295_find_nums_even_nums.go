package main

import (
	"fmt"
)

func main() {
	a := []int{12, 345, 2, 6, 7896}
	fmt.Println(findNumbers(a))
}

func findNumbers(nums []int) int {
	count := 0

	for _, n := range nums {
		if countDigits(n)%2 == 0 {
			count++
		}
	}

	return count
}

func countDigits(n int) int {
	count := 0
	for {
		if n == 0 {
			break
		}
		count++
		n /= 10
	}
	return count
}
