package main

import (
	"fmt"
)

func main() {
	// a := "RLRRLLRLRL"
	// a := "RLLLLRRRLR"
	// a := "LLLLRRRR"
	a := "RLRRRLLRLL"
	fmt.Println(balancedStringSplit(a))
}

func balancedStringSplit(s string) int {
	res := 0
	f := 0

	for i, r := range s {
		if r == 'R' {
			f++
		}

		if r == 'L' {
			f--
		}

		if i != 0 && f == 0 {
			res++
		}
	}

	return res
}
