package main

import (
	"fmt"
	"strconv"
	"strings"
)

func SumDigPow(a, b uint64) []uint64 {
	result := []uint64{}

	for i := a; i <= b; i++ {
		n := strconv.FormatUint(i, 10)
		digits := strings.Split(n, "")
		var sum uint64 = 0
		for j, digit := range digits {
			digitAsNumber, err := strconv.Atoi(digit)
			if err != nil {
				panic("invalid digit on input")
			}

			powValue := powInt(digitAsNumber, j+1)
			sum += powValue
		}

		if sum == i {
			result = append(result, i)
		}
	}

	if len(result) < 1 {
		return nil
	}

	return result
}

func powInt(x, y int) uint64 {
	result := uint64(1)

	for i := 0; i < y; i++ {
		result *= uint64(x)
	}

	return result
}

func main() {
	fmt.Printf("%v \n%v \n", SumDigPow(12157692622039623066, 12157692622039625683), []uint64{12157692622039623539})
}
