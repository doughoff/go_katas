package main

func QueueTime(customers []int, n int) int {
	left, tick, cashiers := customers, 0, make([]int, n)
	println("left:", left, "tick:", tick, "cashiers:", cashiers)

	for len(left) > 0 {
		for _, cashier := range cashiers {
			println("cashier:", cashier)
			if cashier == 0 && len(left) > 0 {
				c, newLeft := left[len(left)-1], left[:len(left)-1]
				left = newLeft
				cashier = c
				println("cashier:", cashier)
			} else {
				cashier--
			}
		}
		println("left:", left, "tick:", tick, "cashiers:", cashiers)

		tick++
	}

	return tick
}

func main() {
	println(
		"result:",
		QueueTime([]int{5, 3, 4}, 1),
	)

	println(
		"result:",
		QueueTime([]int{10, 2, 3, 3}, 2),
	)

	println(
		"result:",
		QueueTime([]int{2, 3, 10}, 2),
	)
}
