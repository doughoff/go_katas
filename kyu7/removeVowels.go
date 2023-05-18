package kata

import "strings"

func Disemvowel(comment string) string {
	vowels := []string{"A", "E", "I", "O", "U"}
	result := ""

	for _, rune := range comment {
		char := string(rune)
		isVowel := false

		for _, vowelRune := range vowels {
			vowelChar := string(vowelRune)
			if vowelChar == strings.ToUpper(char) {
				isVowel = true
			}
		}

		if !isVowel {
			result += char
		}
	}

	return result
}
