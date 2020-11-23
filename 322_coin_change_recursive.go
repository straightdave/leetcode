package main

import "fmt"

func main() {
	coins := []int{2}
	amount := 3
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	m := make(map[int]int)
	return change(m, coins, amount)
}

func min(numbers ...int) int {
	if len(numbers) == 0 {
		panic("numbers is empty")
	}

	if len(numbers) == 1 {
		return numbers[0]
	}

	m := numbers[0]
	for _, n := range numbers[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

func change(m map[int]int, coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	if amount < 0 {
		return -1
	}

	var res []int
	for _, c := range coins {
		var r int
		if v, ok := m[amount-c]; ok {
			r = v
		} else {
			r = change(m, coins, amount-c)
			m[amount-c] = r
		}
		if r == -1 {
			continue
		}
		res = append(res, r)
	}
	if len(res) > 0 {
		return min(res...) + 1
	}
	return -1
}
