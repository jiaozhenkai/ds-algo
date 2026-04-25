package main

import "fmt"

// 参考：https://leetcode.cn/problems/spiral-matrix/solutions/2362055/54-luo-xuan-ju-zhen-mo-ni-qing-xi-tu-jie-juvi/?envType=study-plan-v2&envId=top-100-liked

func spiralOrder(matrix [][]int) []int {
	ans := []int{}
	if len(matrix) == 1 {
		return matrix[0]
	}
	if len(matrix) == 0 {
		return ans
	}

	left := 0                   // 左边界
	right := len(matrix[0]) - 1 // 右边界
	top := 0                    // 上边界
	bottom := len(matrix) - 1   // 下边界
	for {
		// left to right
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		top++
		if top > bottom {
			break
		}

		// top to bottom
		for i := top; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		if right < left {
			break
		}

		// right to left
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[bottom][i])
		}
		bottom--
		if bottom < top {
			break
		}

		// bottom to top
		for i := bottom; i >= top; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left > right {
			break
		}

	}

	return ans
}

func main() {
	matrix := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(spiralOrder(matrix))
}
