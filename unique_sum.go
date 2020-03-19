package main

import (
	"fmt"
)

var powmap = make(map[int]int)

func main() {
	a := []int{2, 2, 3}

	fmt.Println(calc(a))
}

func calc(a []int) int {
	var sum int

	for i, n := range a {
		t := n * pow2(i)
		sum += t
	}

	return sum
}

func pow2(x int) int {
	if n, ok := powmap[x]; ok {
		return n
	}

	if x == 0 {
		return 1
	}

	p := 1
	for i := 0; i < x; i++ {
		p *= 2
	}

	powmap[x] = p
	return p
}
