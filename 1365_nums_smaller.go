package main

import (
	"fmt"
)

func main() {
	// a := []int{6, 5, 4, 8}
	a := []int{7, 7, 7, 7}
	fmt.Println(smallerNumbersThanCurrent(a))
}

func smallerNumbersThanCurrent(nums []int) []int {
	var (
		res   []int
		nums0 = make([]int, len(nums))
	)

	copy(nums0, nums)

	sorted := qsort(nums)

	m := make(map[int]int)

	for i, n := range sorted {
		if _, ok := m[n]; !ok {
			m[n] = i
		}
	}

	for _, n := range nums0 {
		res = append(res, m[n])
	}

	return res
}

func qsort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	pivot := nums[0]

	var (
		smaller []int
		greater []int
	)

	for _, n := range nums[1:] {
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
