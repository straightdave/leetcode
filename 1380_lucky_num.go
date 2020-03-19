package main

import (
	"fmt"
)

func main() {
	m := [][]int{
		[]int{3, 7, 8},
		[]int{9, 11, 13},
		[]int{15, 16, 17},
	}
	fmt.Println(luckyNumbers(m))
}

func luckyNumbers(matrix [][]int) []int {
	var res []int

	if len(matrix) == 0 {
		return res
	}

	var (
		minOfRows []int
		maxOfCols []int
	)

	for _, row := range matrix {
		minOfRows = append(minOfRows, findMinIndex(row))
	}

	for c := 0; c < len(matrix[0]); c++ {
		var col []int
		for r := 0; r < len(matrix); r++ {
			col = append(col, matrix[r][c])
		}
		maxOfCols = append(maxOfCols, findMaxIndex(col))
	}

	for i, colIndex := range minOfRows {
		if maxOfCols[colIndex] == i {
			res = append(res, matrix[i][colIndex])
		}
	}
	return res
}

func findMinIndex(a []int) int {
	minIndex := 0
	min := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
			minIndex = i
		}
	}
	return minIndex
}

func findMaxIndex(a []int) int {
	maxIndex := 0
	max := a[0]

	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
			maxIndex = i
		}
	}
	return maxIndex
}
