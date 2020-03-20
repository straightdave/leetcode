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

//
// NOTICE: this works but will fail by "Time Limit Exceeds"
//
func reconstructQueue(people [][]int) [][]int {
	people = qsort(people)

	index := 0
	for {
		if index == len(people) {
			// Done!
			break
		}

		if ok, greaterBefore := needsMove(people, index); ok {
			people = move(people, index, greaterBefore) // would change people!
			// index = 0
		} else {
			index++
		}
	}

	return people
}

func move(people [][]int, index, greaterBefore int) [][]int {
	ph := people[index][0]
	pk := people[index][1]
	howManyGreaterNeedsToSurpass := pk - greaterBefore

	// calculate the newIndex
	newIndex := -1
	for i := index + 1; i < len(people); i++ {
		p := people[i]

		if p[0] >= ph {
			howManyGreaterNeedsToSurpass--
			if howManyGreaterNeedsToSurpass == 0 {
				newIndex = i + 1
				break
			}
		}
	}

	// if newIndex's valid
	if newIndex > index {
		// insert one copy to newIndex
		cp := make([]int, 2)
		copy(cp, people[index])

		// copy tail
		tail := people[newIndex:]
		tail2 := make([][]int, len(tail))
		copy(tail2, tail)

		people = append(append(people[:newIndex], cp), tail2...)

		// remove old one from index
		people = append(people[:index], people[index+1:]...)
	}

	return people
}

func needsMove(people [][]int, index int) (bool, int) {
	ph := people[index][0]
	pk := people[index][1]
	count := 0
	for _, p := range people[:index] {
		if p[0] >= ph {
			count++
		}
	}
	return pk != count, count
}

// sort by h (first) as well as k (second)
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
		} else if p[0] == pivot[0] && p[1] < pivot[1] {
			smaller = append(smaller, p)
		} else {
			greater = append(greater, p)
		}
	}

	smaller = qsort(smaller)
	greater = qsort(greater)
	return append(append(smaller, pivot), greater...)
}
