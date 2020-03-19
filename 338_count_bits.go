package main

import (
    "fmt"
)

func main() {
    fmt.Println(countBits(5))
}

func countBits(num int) []int {
    var r []int
    for i := 0; i <= num; i++ {
        r = append(r, count(i))
    }
    return r
}

func count(n int) int {
    c := 0

    for n != 0 {
        if n%2 == 1 {
            c++
        }
        n /= 2
    }

    return c
}
