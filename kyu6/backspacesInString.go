package main

import (
	"fmt"
	"strings"
)

func CleanString(s string) string {
	chars := strings.Split(s, "")
	res := []string{}

	for i := 0; i < len(chars); i++ {
		char := chars[i]
		if char == "#" {
			if len(res) > 0 {
				res = res[:len(res)-1]
			}

			continue
		}

		res = append(res, char)
	}

	return strings.Join(res, "")
}

func main() {
	fmt.Println(CleanString("abc#d##c"), "||", "ac")
	fmt.Println(CleanString("abc####d##c#"), "||", "")
	fmt.Println(CleanString(""), "||", "")
	fmt.Println(CleanString("#######"), "||", "")
}
