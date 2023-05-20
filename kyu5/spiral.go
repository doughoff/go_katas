package main

import "fmt"

func CreateSpiral(n int) [][]int {
	direction, x, y := "r", 0, 0

	result := make([][]int, n)

	for i := range result {
		result[i] = make([]int, n)
	}

	for curr := 1; curr <= n*n; curr++ {
		result[y][x] = curr
		if direction == "r" {

			if x+1 >= n || result[y][x+1] != 0 {
				direction = "d"
				y++
				continue
			}

			x++

		} else if direction == "l" {

			if x-1 < 0 || result[y][x-1] != 0 {
				direction = "u"
				y--
				continue
			}

			x--

		} else if direction == "d" {

			if y+1 >= n || result[y+1][x] != 0 {
				direction = "l"
				x--
				continue
			}

			y++

		} else if direction == "u" {

			if y-1 < 0 || result[y-1][x] != 0 {
				direction = "r"
				x++
				continue
			}

			y--
		}
	}

	return result
}

func main() {
	var (
		oneByOne     = [][]int{{1}}
		twoByTwo     = [][]int{{1, 2}, {4, 3}}
		threeByThree = [][]int{{1, 2, 3}, {8, 9, 4}, {7, 6, 5}}
	)

	fmt.Printf("%v || %v\n", CreateSpiral(1), oneByOne)
	fmt.Printf("%v || %v\n", CreateSpiral(2), twoByTwo)
	fmt.Printf("%v || %v\n", CreateSpiral(3), threeByThree)
}
