package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3, 4}

	fmt.Println(append(a[:2], 1))

	// a has been changed
	fmt.Println(a)

	a = append(append(a[:2], 1), a[2:]...)

	fmt.Println(a)

}
