package main

import (
	"fmt"
)

func main() {
	a := [][]int{
		[]int{7, 0},
		[]int{4, 4},
		[]int{7, 1},
		[]int{5, 0},
		[]int{6, 1},
		[]int{5, 2},
	}
	fmt.Println(reconstructQueue(a))
}

func reconstructQueue(people [][]int) [][]int {
	people = qsort(people)

	// as in such great order, we can just move each to its 'k' (as the new index)
	for i := 0; i < len(people); i++ {
		people = move(people, i)
	}
	return people
}

func move(people [][]int, index int) [][]int {
	// move people[index] to people[ k ]
	k := people[index][1]

	if k == index {
		// no need to move
		return people
	}

	// remove it from people first
	p := make([]int, 2)
	copy(p, people[index])

	people = append(people[:index], people[index+1:]...)

	// insert it to new people[k]
	tail := make([][]int, len(people[k:]))
	copy(tail, people[k:])

	people = append(append(people[:k], p), tail...)
	return people
}

// qsort by h (first) as well as k (second)
// greater -> smaller (but for k, smaller -> greater)
func qsort(a [][]int) [][]int {
	if len(a) < 2 {
		return a
	}

	var (
		smaller [][]int
		greater [][]int
	)

	pivot := a[0]
	for _, p := range a[1:] {
		if p[0] < pivot[0] {
			smaller = append(smaller, p)
		} else if p[0] == pivot[0] {
			if p[1] < pivot[1] {
				greater = append(greater, p)
			} else {
				smaller = append(smaller, p)
			}
		} else {
			greater = append(greater, p)
		}
	}

	smaller = qsort(smaller)
	greater = qsort(greater)
	return append(append(greater, pivot), smaller...)
}
