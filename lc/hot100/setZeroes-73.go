package main

/*
思路：优化空间复杂度为 O(1)。。。。。。
使用第一行第一列对应位置作为数组保存本行本列是否有 0;因为第一行第一列可能被更改，所以先保存看第一行第一列是否有 0；
- 如果有0，后续直接第一行第一列全部设置为0即可,
- 如果没有0，不用管了，其他行列为如果有 0 会把第一行第一列对应位置改为 0 的。

如果第一行有 0，比如[0,4] 位置是 0，[i,4] 的其他位置都没有 0 ，但是在判断时，发现[0,4]原来就是 0，[i,4] 也都会被设置为 0

没啥意义的优化。。。
*/

func setZeroes1(matrix [][]int) {
	isFirstRowHasZero := false
	isFirstColHasZero := false
	for i := 0; i < len(matrix[0]); i++ {
		if matrix[0][i] == 0 {
			isFirstRowHasZero = true
			break
		}
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == 0 {
			isFirstColHasZero = true
			break
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				matrix[0][j] = 0
				matrix[i][0] = 0
			}
		}
	}

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if matrix[0][j] == 0 || matrix[i][0] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	// 第一行全部设置为 0
	if isFirstRowHasZero {
		for i := 0; i < len(matrix[0]); i++ {
			matrix[0][i] = 0
		}
	}
	// 第一列全部设置为 0
	if isFirstColHasZero {
		for i := 0; i < len(matrix); i++ {
			matrix[i][0] = 0
		}
	}
}

/*
思路：先遍历，hash 记录出现 0 的位置。再次遍历 hash, 将出现 0 位置的行列全部设置为 0；
空间复杂度为 O(n)
*/
func setZeroes(matrix [][]int) {
	type Flag struct {
		i int
		j int
	}
	hash := map[Flag]struct{}{}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				flag := Flag{i: i, j: j}
				hash[flag] = struct{}{}
			}
		}
	}

	for k, _ := range hash {
		row := k.i
		col := k.j
		// 第 i 列全部变为 0
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
		// 第 j 行全部变为 0
		for j := 0; j < len(matrix[0]); j++ {
			matrix[row][j] = 0
		}
	}
}
