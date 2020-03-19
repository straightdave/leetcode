package main

import (
	"fmt"
)

func main() {
	fmt.Println(subtractProductAndSum(4421))
}

func subtractProductAndSum(n int) int {
	var (
		digits []int
		prod   = 1
		sum    = 0
	)

	for {
		if n == 0 {
			break
		}

		digits = append(digits, n%10)
		n /= 10
	}

	for _, n := range digits {
		prod *= n
		sum += n
	}

	return prod - sum
}
