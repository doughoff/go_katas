package main

import "strings"

func AbbrevName(name string) string {
	names := strings.Split(name, " ")
	result := []string{}

	for _, s := range names {
		letters := strings.Split(s, "")
		result = append(result, strings.ToUpper(letters[0]))
	}

	return strings.Join(result, ".")
}
