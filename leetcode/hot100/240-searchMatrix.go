package main

func main() {
	//m := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	//searchMatrix2(m, 20)

	m1 := [][]int{{-5}}
	searchMatrix2(m1, -10)

}

/*
以右上角的元素为锚点 p
if p > target, 此时 p 所在的这一列可以被消掉，左移一列
if p < target, 此时 p 所在的这一行可以被消掉，下移一行
整体趋势是在往左下角走

本质上，比较大小有传递性，找到一个点，这个点的大小可以通过传递性抹去一些无关的细节，用来降低时间复杂度，比如：
a < b < c; 如果 d > c 那么 a,b 一定小于 d; 如果 d < a, 那么 b,c 一定大于 d
*/
func searchMatrix2(matrix [][]int, target int) bool {
	row := len(matrix)
	col := len(matrix[0])

	i := 0
	j := col - 1

	for i < row && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else if matrix[i][j] < target {
			i++
		}
	}
	return false
}

/*
非矩阵
第一行从单调递增，以第一行每一个元素开始的列都是单调递增的; m[0][j] 这一行左边的行一定小于右边的行。那么可以：
if m[0][j] == t; return true
if m[0][j] > t; return false（后面的列不用搜索了）
if m[0][j] < t && m[n-1][j] > t; 可能在这一列
  - 其他情况继续搜索
*/
func searchMatrix1(matrix [][]int, target int) bool {
	n := len(matrix[0]) // col
	m := len(matrix)    // row
	for i := 0; i < n; i++ {
		if matrix[0][i] == target {
			return true
		}

		if matrix[0][i] > target {
			return false
		}

		if matrix[0][i] < target && matrix[m-1][i] >= target {
			for j := 1; j < m; j++ {
				if matrix[j][i] == target {
					return true
				}
			}
		}
	}

	return false
}

/*
如果是矩阵可以按照下面的方法来。

主对角线将 matrix 分为两部分，左下部分的数值一定小于右上部分；并且 matrix[i][j] 一定是 matrix[0..i][0..j] 的最大值

if matrix[i][j] == target return true
if matrix[i][j] < target
- matrix[0..i][0..j] 这部分分块矩阵的值一定都是小于 target 的，需要继续沿着主对角线继续
if matrix[i][j] > target
- 沿着第 i 行，第 j 列搜索是否包含 target
	- 如果有 return ture
	- 如果没有，继续沿着主对角线继续
*/

func searchMatrix(matrix [][]int, target int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		if matrix[i][i] == target {
			return true
		}
		if matrix[i][i] < target {
			continue
		}
		if matrix[i][i] > target {
			bRow := false
			bCol := false
			for j := 0; j < i; j++ {
				if matrix[j][i] == target {
					bRow = true
				}
				if matrix[i][j] == target {
					bCol = true
				}
			}
			if bRow || bCol {
				return true
			}
		}
	}

	return false
}
