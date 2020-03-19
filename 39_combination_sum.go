package main

import (
	"fmt"
)

func main() {
	a := []int{2, 3, 6, 7}
	target := 7

	// a := []int{2, 3, 5}
	// target := 8
	fmt.Println(combinationSum(a, target))
}

func combinationSum(candidates []int, target int) [][]int {
	candidates = qsort(candidates)
	// fmt.Println("sorted:", candidates)

	lenCandidate := len(candidates)

	var stackOfIndex []int
	var sumInStack int
	var res [][]int

	f := func(target int) {
		index := 0
		for {
			// whenever "index to push" exceeds,
			// we need to try larger one in previous position
			for index >= lenCandidate {
				// cannot try previous, quit
				if len(stackOfIndex) < 1 {
					return
				}

				// set "index to push" to be "previous one +1"
				index = stackOfIndex[len(stackOfIndex)-1] + 1

				// pop previous and wait to push new index to previous position
				sumInStack -= candidates[stackOfIndex[len(stackOfIndex)-1]]
				stackOfIndex = stackOfIndex[:len(stackOfIndex)-1]
			}

			// push one
			stackOfIndex = append(stackOfIndex, index)
			sumInStack += candidates[index]

			if sumInStack < target {
				// just push new index to the stack
				// only less, extend the stack
				continue
			}

			// below is "current sum equals or exceeds target"

			if sumInStack == target {
				// one success
				var res1 []int
				for _, i := range stackOfIndex {
					res1 = append(res1, candidates[i])
				}
				res = append(res, res1)
			}

			if len(stackOfIndex) == 1 {
				// sum >= target, and it's the only one in stack,
				// quit
				return
			}

			// pop current
			sumInStack -= candidates[stackOfIndex[len(stackOfIndex)-1]]
			stackOfIndex = stackOfIndex[:len(stackOfIndex)-1]

			// peek previous, and set to "+1"
			index = stackOfIndex[len(stackOfIndex)-1] + 1

			// pop previous and so to push new index to previous position
			sumInStack -= candidates[stackOfIndex[len(stackOfIndex)-1]]
			stackOfIndex = stackOfIndex[:len(stackOfIndex)-1]
		}
	}

	f(target)
	return res
}

func qsort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	var (
		smaller []int
		greater []int
	)

	pivot := a[0]
	for _, n := range a[1:] {
		if n > pivot {
			greater = append(greater, n)
		} else {
			smaller = append(smaller, n)
		}
	}

	smaller = qsort(smaller)
	greater = qsort(greater)
	return append(append(smaller, pivot), greater...)
}
