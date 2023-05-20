package main

import (
	"fmt"
	"strconv"
	"strings"
)

func ValidISBN10(isbn string) bool {
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for pos, rune := range isbn {
		number, err := strconv.Atoi(string(rune))

		if err != nil && pos != 9 {
			return false
		}

		if err != nil {
			if strings.ToLower(string(rune)) != "x" {
				return false
			}
			number = 10
		}

		sum += number * (pos + 1)
	}

	return sum%11 == 0
}

func main() {
	fmt.Print(ValidISBN10("1112223339"), " ", "1112223339 is valid\n")
	fmt.Print(ValidISBN10("048665088X"), " ", "048665088X is valid\n")
	fmt.Print(ValidISBN10("1293000000"), " ", "1293000000 is valid\n")
	fmt.Print(ValidISBN10("1234554321"), " ", "1234554321 is valid\n\n")

	fmt.Print(ValidISBN10("1234512345"), " ", "1234512345 is NOT valid\n")
	fmt.Print(ValidISBN10("1293"), " ", "1293 is NOT valid\n")
	fmt.Print(ValidISBN10("X123456788"), " ", "X123456788 is NOT valid\n")
	fmt.Print(ValidISBN10("ABCDEFGHIJ"), " ", "ABCDEFGHIJ is NOT valid\n")
	fmt.Print(ValidISBN10("XXXXXXXXXX"), " ", "XXXXXXXXXX is NOT valid\n")
}
