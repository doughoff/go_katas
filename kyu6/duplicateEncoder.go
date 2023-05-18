package main

import (
	"fmt"
	"strings"
)

func DuplicateEncode(word string) string {
	charMap := make(map[string]int)

	charArray := strings.Split(word, "")

	for _, char := range charArray {
		charMap[strings.ToLower(char)]++
	}

	result := []string{}

	for _, char := range charArray {
		if charMap[strings.ToLower(char)] > 1 {
			result = append(result, ")")
		} else {
			result = append(result, "(")
		}
	}

	return strings.Join(result, "")
}

func main() {
	fmt.Println(DuplicateEncode("din"), "||", "(((")
	fmt.Println(DuplicateEncode("recede"), "||", "()()()")
	fmt.Println(DuplicateEncode("(( @"), "||", "))((")
}
