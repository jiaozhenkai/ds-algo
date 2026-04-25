package main

import (
	"fmt"
	"slices"
)

/*
先按照主对角线反转，再按照行反转
1 2 3      1 4 7     7 4 1
4 5 6  -》 2 5 8  -》 8 5 2
7 8 9      3 6 9     9 6 3
*/

func rotate48(matrix [][]int) {
	// 主对角线反转,
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 按照行反转
	for i := 0; i < n; i++ {
		slices.Reverse(matrix[i])
	}
}

func main() {
	m := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate48(m)

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m); j++ {
			fmt.Printf("%d ", m[i][j])
		}
		fmt.Println()
	}

}
