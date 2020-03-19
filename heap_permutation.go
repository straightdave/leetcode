package main

import (
    "fmt"
)

func main() {
    a := []int{1, 2, 3}

    var res [][]int
    var gen func(int)

    gen = func(n int) {
        if n == 1 {
            aa := make([]int, len(a))
            copy(aa, a)
            res = append(res, aa)
            return
        }

        for i := 0; i < n; i++ {
            gen(n - 1)

            if n % 2 == 1 {
                swap(a, 0, n - 1)
            } else {
                swap(a, i, n - 1)
            }
        }
    }

    gen(len(a))
    fmt.Println(res)
}

func swap(a []int, x, y int) {
    a[x], a[y] = a[y], a[x]
}

