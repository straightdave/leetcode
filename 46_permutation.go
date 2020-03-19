package main

import (
    "fmt"
)

func main() {
    a := []int{1, 2, 3}
    fmt.Println(permute(a))
}

func permute(nums []int) [][]int {
    // Heap's permutation algorithm
    // dave.wu 2020-03-10

    var res [][]int
    var gen func(int)

    gen = func(n int) {
        if n == 1 {
            aa := make([]int, len(nums))
            copy(aa, nums)
            res = append(res, aa)
            return
        }

        for i := 0; i < n; i++ {
            gen(n - 1)
            if n%2 == 1 {
                swap(nums, 0, n-1)
            } else {
                swap(nums, i, n-1)
            }
        }
    }

    gen(len(nums))
    return res
}

func swap(a []int, x, y int) {
    a[x], a[y] = a[y], a[x]
}
