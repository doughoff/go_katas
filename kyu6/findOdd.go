package main

import "fmt"

func FindOdd(seq []int) int {
	numberMap := make(map[int]int)
	result := 0

	for _, n := range seq {
		numberMap[n]++
	}

	for key, value := range numberMap {
		if value%2 != 0 {
			result = key
			break
		}
	}

	return result
}

func main() {
	arr := [...][]int{
		{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5},
		{1, 1, 2, -2, 5, 2, 4, 4, -1, -2, 5},
		{20, 1, 1, 2, 2, 3, 3, 5, 5, 4, 20, 4, 5},
		{10},
		{1, 1, 1, 1, 1, 1, 10, 1, 1, 1, 1},
		{5, 4, 3, 2, 1, 5, 4, 3, 2, 10, 10},
	}
	sol := [...]int{5, -1, 5, 10, 10, 1}
	for i, v := range arr {
		fmt.Println(FindOdd(v), "||", sol[i])
	}
}
