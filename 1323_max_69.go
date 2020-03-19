package main

import (
	"fmt"
)

func main() {
	// fmt.Println(maximum69Number(9669))
	fmt.Println(maximum69Number(9999))
}

func maximum69Number(num int) int {
	var digits []int
	x := num
	for {
		if x == 0 {
			break
		}
		digits = append(digits, x%10)
		x /= 10
	}

	for i := len(digits) - 1; i >= 0; i-- {
		if digits[i] == 6 {
			digits[i] = 9
			break
		}
	}

	res := 0
	for i := len(digits) - 1; i >= 0; i-- {
		res = res*10 + digits[i]
	}
	return res
}
