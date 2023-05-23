package main

func Cakes(recipe, available map[string]int) int {
	var min int
	started := false

	for k, v := range recipe {
		canProduce := available[k] / v
		if !started || min > canProduce {
			min = canProduce
			started = true
		}
	}

	return min
}
