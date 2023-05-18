package kata

import "strings"

func Greet(name string) string {
	// fill in solution here
	return strings.Replace("Hello, <name> how are you doing today?", "<name>", name, 1)
}
