package main

import (
	"fmt"
	"strings"
)

func TowerBuilder(nFloors int) []string {
	tree := []string{}

	for i := 0; i < nFloors; i++ {
		floor := ""

		for j := 0; j <= i; j++ {
			floor += "*"
		}

		if len(floor) < nFloors {
			floor += strings.Repeat(" ", nFloors-len(floor))
		}

		tree = append(tree, floor)
	}

	mirroredTree := []string{}

	for i := 0; i < nFloors; i++ {
		temp := ""
		for j := len(tree[i]) - 1; j > 0; j-- {
			temp += string(tree[i][j])
		}

		mirroredTree = append(mirroredTree, temp)
	}

	for i := 0; i < nFloors; i++ {
		tree[i] = mirroredTree[i] + tree[i]
	}

	return tree
}

func main() {
	result := TowerBuilder(5)

	for _, floor := range result {
		fmt.Println(floor)
	}
}
