package main

import (
	"fmt"
)

func main() {
	a := []int{17, 18, 5, 4, 6, 1}
	fmt.Println(replaceElements(a))
}

func replaceElements(arr []int) []int {
	var max = make([]int, len(arr))

	for i := len(arr) - 1; i >= 0; i-- {
		if i == len(arr)-1 {
			max[i] = -1
			continue
		}

		if i == len(arr)-2 {
			max[i] = arr[i+1]
			continue
		}

		if arr[i+1] > max[i+1] {
			max[i] = arr[i+1]
		} else {
			max[i] = max[i+1]
		}
	}

	return max
}
