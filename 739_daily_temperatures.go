package main

import (
	"fmt"
)

func main() {
	a := []int{55,38,53,81,61,93,97,32,43,78}
	b := dailyTemperatures(a)
	fmt.Println(b)
}

func dailyTemperatures(T []int) []int {
	n := len(T)
	var res []int

	for i := 0; i < n - 1; i++ {
		j := i + 1

		x := 0
		for j < n {
			if T[i] < T[j] {
				break
			}
			j++
		}

		if j == n {
			res = append(res, 0)
		} else {
			res = append(res, j-i)
		}
	}

	res = append(res, 0)
	return res
}
