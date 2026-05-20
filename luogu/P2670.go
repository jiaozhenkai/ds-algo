package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
两种思路：
1. 读取输入到二维数组，然后遍历二维数组，如果当前是 * ，打印* ，如果不是 * ，看各个方位的值，但是这样要考虑好边界条件，可以使用虚拟边界来简化运算。。
2. 反向考虑，遇到 * ，就看记录他的 8 个方向，每个方向的横纵坐标为 key, value 是炸弹的数量，每次遇到一个 * 这 8 个方向一定炸弹数加 1 ；后续通过构造正确的横纵坐标作为 key 去 map 中获取 value，这样不用考虑边界条件了。这种方式还有一个点要考虑：遇到 * 要记录好当前 * 的坐标，以便于后续对比打印，可以用一个新的 map 来保存

这两种思路本质一个是从本节点考虑对应的 8 个方向是否有炸弹来更新自己，一个是从本节点考虑对应的 8 个方向是否有炸弹来更新这 8 个方向。
*/

// 思路二
func main1() {
	type pair struct {
		i int
		j int
	}

	hash := make(map[pair]int)
	hash1 := make(map[pair]struct{})

	f := func(p pair) {
		if _, ok := hash[p]; ok {
			hash[p] += 1
		} else {
			hash[p] = 1
		}
	}

	var n, m int
	fmt.Scan(&n, &m)

	col, row := 0, 0

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		for _, v := range s {
			if v == '*' {
				hash1[pair{i: row, j: col}] = struct{}{}
				//左上
				leftTop := pair{i: row - 1, j: col - 1}
				f(leftTop)
				// 左下
				leftBottom := pair{i: row + 1, j: col - 1}
				f(leftBottom)
				//左
				left := pair{i: row, j: col - 1}
				f(left)
				// 上
				top := pair{i: row - 1, j: col}
				f(top)
				// 下
				bottom := pair{i: row + 1, j: col}
				f(bottom)
				// 右
				right := pair{i: row, j: col + 1}
				f(right)
				// 右上
				rightTop := pair{i: row - 1, j: col + 1}
				f(rightTop)
				// 右下
				rightBottom := pair{i: row + 1, j: col + 1}
				f(rightBottom)
			}
			col++
		}
		col = 0
		row++
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			p := pair{i: i, j: j}
			if _, ok := hash1[p]; ok {
				fmt.Printf("%s", "*")
				continue
			}
			if hash[p] == 0 {
				fmt.Printf("%d", 0)
			} else {
				fmt.Printf("%d", hash[p])
			}
		}
		fmt.Println()
	}
}

// 思路一
func main() {
	// 上下左右，左上，左下，右上，右下
	dx := []int{-1, 1, 0, 0, -1, 1, -1, 1}
	dy := []int{0, 0, -1, 1, -1, -1, 1, 1}
	var n, m int
	fmt.Scan(&n, &m)
	scanner := bufio.NewScanner(os.Stdin)
	arr := [102][102]byte{}
	for i := 1; i <= n; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j := 1; j <= m; j++ {
			arr[i][j] = line[j-1]
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if arr[i][j] == '*' {
				fmt.Printf("%c", arr[i][j])
			} else {
				cnt := 0
				for k := 0; k < 8; k++ {
					if arr[i+dx[k]][j+dy[k]] == '*' {
						cnt++
					}
				}
				fmt.Printf("%d", cnt)
			}
		}
		fmt.Println()
	}
}
