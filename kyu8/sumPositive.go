package kata

func PositiveSum(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		if num > 0 {
			sum += num
		}
	}

	return sum
}
