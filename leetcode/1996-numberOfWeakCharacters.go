package main

import (
	"fmt"
	"sort"
)

/*
 */

func numberOfWeakCharacters(properties [][]int) int {

	sort.Slice(properties, func(i, j int) bool {
		if properties[i][0] != properties[j][0] {
			return properties[i][0] > properties[j][0]
		}
		return properties[i][1] < properties[j][1]

	})

	count := 0
	maxD := properties[0][1]

	for i := 1; i < len(properties); i++ {
		if properties[i][1] < maxD {
			count++
		} else {
			maxD = properties[i][1]
		}
	}

	return count
}

func main() {
	p := [][]int{{5, 5}, {6, 3}, {3, 6}}
	res := numberOfWeakCharacters(p)
	fmt.Println(res)

	p1 := [][]int{{1, 5}, {10, 4}, {4, 3}}
	res1 := numberOfWeakCharacters(p1)
	fmt.Println(res1)
}
