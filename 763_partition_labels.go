package main

import (
	"fmt"
)

func partitionLabels(S string) []int {
	var res []int

	for i := 0; i < len(S); i++ {
		end := farmost(i, S)

		if i <= end {
			res = append(res, end-i+1)
		}

		i = end
	}

	return res
}

func farmost(startIndex int, S string) int {

	ending := right(S[startIndex], S)
	if ending <= startIndex+1 {
		return ending
	}

	for i := startIndex + 1; i <= ending; i++ {
		end := right(S[i], S)
		if end > ending {
			ending = end
		}
	}

	return ending
}

func right(c byte, S string) int {
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == c {
			return i
		}
	}
	return -1
}

func main() {
	// s := "qiejxqfnqceocmy"
	s := "ababcbacadefegdehijhklij"
	fmt.Printf("len=%d: %+v\n", len(s), partitionLabels(s))
}
