// https://www.codewars.com/kata/52597aa56021e91c93000cb0/train/go
package main

import "fmt"

func MoveZeros(arr []int) []int {
	zeroes := []int{}
	numbers := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			zeroes = append(zeroes, 0)
		} else {
			numbers = append(numbers, arr[i])
		}
	}

	return append(numbers, zeroes...)
}

func main() {
	fmt.Printf("%v", MoveZeros([]int{1, 2, 0, 3, 5, 0, 6, 0}))
}
