package kata

import "strings"

type MyString string

func (s MyString) IsUpperCase() bool {
	for _, rune := range string(s) {
		char := string(rune)
		if char != strings.ToUpper(char) {
			return false
		}
	}

	return true
}
