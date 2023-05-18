package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func SumStrNumbers(str string) int {
	chars := strings.Split(str, "")
	sum := 0

	for i := 0; i < len(chars); i++ {
		n, err := strconv.Atoi(chars[i])
		if err != nil {
			panic("invalid a number")
		}

		sum += n
	}

	return sum
}

func OrderWeight(str string) string {
	tokens := strings.Split(str, " ")

	sort.Slice(tokens, func(i, j int) bool {
		a := SumStrNumbers(tokens[i])
		b := SumStrNumbers(tokens[j])

		if a != b {
			return a < b
		}

		tempArr := []string{tokens[i], tokens[j]}
		sort.Strings(tempArr)

		if tempArr[0] == tokens[i] {
			return true
		}

		return false

	})

	return strings.Join(tokens, " ")
}

func main() {
	fmt.Println(OrderWeight("103 123 4444 99 2000"), "||", "2000 103 123 4444 99")
	fmt.Println(OrderWeight("2000 10003 1234000 44444444 9999 11 11 22 123"), "||", "11 11 2000 10003 22 123 1234000 44444444 9999")
	fmt.Println(OrderWeight(""), "||", "")
}
