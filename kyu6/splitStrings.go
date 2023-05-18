package kata

func Solution(str string) []string {
	lag, result := "", []string{}
	for _, rune := range str {
		char := string(rune)
		if len(lag) == 0 {
			lag = char
		} else {
			result = append(result, lag+char)
			lag = ""
		}
	}

	if len(lag) > 0 {
		result = append(result, lag+"_")
	}

	return result
}
