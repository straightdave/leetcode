package main

import (
	"fmt"
)

func main() {
	fmt.Println(numberOfSteps(8))
}

func numberOfSteps(num int) int {
	var (
		m = make(map[int]int)
		f func(int) int
	)

	f = func(n int) int {
		if s, ok := m[n]; ok {
			return s
		}

		if n == 0 {
			return 0
		}

		var n2 int
		if n%2 == 0 {
			n2 = n / 2
		} else {
			n2 = n - 1
		}

		s := f(n2)
		m[n] = s + 1
		return s + 1
	}

	return f(num)
}
