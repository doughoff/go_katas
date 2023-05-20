package main

import (
	"fmt"
	"strings"
)

func FirstNonRepeating(str string) string {
	charMap := make(map[string]int)

	charArray := strings.Split(str, "")
	for _, char := range charArray {
		charMap[strings.ToLower(char)]++
	}

	result := ""

	for _, char := range charArray {
		if charMap[strings.ToLower(char)] == 1 {
			result = char
			break
		}
	}

	//your code here
	return result
}

func main() {

	fmt.Println(FirstNonRepeating("a"), "||", "a")
	fmt.Println(FirstNonRepeating("stress"), "||", "t")
	fmt.Println(FirstNonRepeating("moonmen"), "||", "e")

	fmt.Println(FirstNonRepeating(""), "||", "")

	fmt.Println(FirstNonRepeating("abba"), "||", "")
	fmt.Println(FirstNonRepeating("aa"), "||", "")

	fmt.Println(FirstNonRepeating("~><#~><"), "||", "#")
	fmt.Println(FirstNonRepeating("hello world, eh?"), "||", "w")

	fmt.Println(FirstNonRepeating("sTreSS"), "||", "T")
	fmt.Println(FirstNonRepeating("Go hang a salami, I'm a lasagna hog!"), "||", ",")
}
