package main

import (
	"fmt"
)

func main() {
	a := []int{1, 1, 2, 3}
	fmt.Println(decompressRLElist(a))
}

func decompressRLElist(nums []int) []int {
	var (
		order  []int
		orderV []int // tricky: don't use map to store times, value may duplicate
		res    []int
	)

	for i := 0; i < len(nums); i += 2 {
		order = append(order, nums[i+1])
		orderV = append(orderV, nums[i])
	}

	for i, n := range order {
		times := orderV[i]
		for i := 0; i < times; i++ {
			res = append(res, n)
		}
	}

	return res
}
