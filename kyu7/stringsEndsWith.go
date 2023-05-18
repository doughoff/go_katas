package main

func solution(str, ending string) bool {
	if str == ending {
		return true
	}

	result := true

	for i := 1; (i < len(str)) && (i < len(ending)); i++ {
		if str[len(str)-i] != ending[len(ending)-i] {
			result = false
			break
		}
	}

	return result
}
