package main

func QueueTime(customers []int, n int) int {
	var minLen int
	if len(customers) < n {
		minLen = len(customers)
	} else {
		minLen = n
	}

	left, tick, cashiers := customers, 0, make([]int, minLen)

	for len(left) > 0 {
		for i := 0; i < len(cashiers); i++ {
			if cashiers[i] == 0 && len(left) > 0 {
				var c int
				c, left = left[len(left)-1], left[:len(left)-1]
				cashiers[i] = c - 1
			} else {
				cashiers[i] -= 1
			}
		}
		tick++
	}

	for len(cashiers) > 0 {
		hasCustomer := false
		for i := 0; i < len(cashiers); i++ {
			if cashiers[i] == 0 {
				cashiers = append(cashiers[:i], cashiers[i+1:]...)
			} else {
				hasCustomer = true
				cashiers[i] -= 1
			}
		}

		if hasCustomer {
			tick++
		}
	}

	return tick
}

func main() {
	// println(
	// 	"result:",
	// 	QueueTime([]int{5, 3, 4}, 1),
	// )

	// println(
	// 	"result:",
	// 	QueueTime([]int{10, 2, 3, 3}, 2),
	// )

	println(QueueTime([]int{}, 1), 0)
	println(QueueTime([]int{1, 2, 3, 4}, 1), 10)
	println(QueueTime([]int{2, 2, 3, 3, 4, 4}, 2), 9)
	println(QueueTime([]int{1, 2, 3, 4, 5}, 100), 5)

	// println(
	// 	"result:",
	// 	QueueTime([]int{2, 3, 10}, 2),
	// )
}
