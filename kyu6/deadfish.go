package main

import "fmt"

func Parse(data string) []int {
	result := []int{}
	value := 0

	for _, c := range data {
		switch c {
		case 'i':
			value++
		case 'd':
			value--
		case 's':
			value *= value
		case 'o':
			result = append(result, value)
		}
	}

	return result
}

func main() {
	fmt.Printf("%v: %v \n", Parse("ooo"), []int{0, 0, 0})
	fmt.Printf("%v: %v \n", Parse("ioioio"), []int{1, 2, 3})
	fmt.Printf("%v: %v \n", Parse("idoiido"), []int{0, 1})
	fmt.Printf("%v: %v \n", Parse("isoisoiso"), []int{1, 4, 25})
	fmt.Printf("%v: %v \n", Parse("codewars"), []int{0})
}
