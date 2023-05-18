package main

import "fmt"

func MaximumSubarraySum(numbers []int) int {
	result := 0

	for i := 0; i < len(numbers); i++ {
		sum := 0
		for j := i; j < len(numbers); j++ {
			sum += numbers[j]

			if sum > result {
				result = sum
			}
		}
	}

	return result
}

func main() {
	fmt.Printf("%v should be 6 \n", MaximumSubarraySum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))
}
