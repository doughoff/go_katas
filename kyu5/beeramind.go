package main

import "fmt"

// 3,5,7,9 >> +2 in each level
// the sum of all the previous lol
// starting with 1.
// 1,3,5,7,9, 11 ...
// results : 1, 4, 9, 16, 25, 36

func Beeramid(bonus int, price float64) int {
	beerAmount := int(float64(bonus) / price)
	levels, beerUsed, beerUsedInLevel, progression := 1, 1, 1, 3

	if beerAmount < 1 {
		return 0
	}

	for beerUsed < beerAmount {
		beerUsedInLevel += progression
		beerUsed += beerUsedInLevel
		progression += 2
		if beerUsed <= beerAmount {
			levels++
		}
	}

	return levels
}

func main() {
	fmt.Println(Beeramid(9, 2.0), "||", 1)
	fmt.Println(Beeramid(21, 1.5), "||", 3)
	fmt.Println(Beeramid(-1, 4.0), "||", 0)
}
