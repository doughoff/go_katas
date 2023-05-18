package main

func Arithmetic(a int, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	default:
		return 0
	}
}

func main() {
	println(Arithmetic(8, 2, "add"), 10)
	println(Arithmetic(8, 2, "subtract"), 6)
	println(Arithmetic(8, 2, "multiply"), 16)
	println(Arithmetic(8, 2, "divide"), 4)
}
