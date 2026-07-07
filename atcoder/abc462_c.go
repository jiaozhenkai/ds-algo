package main

import (
	"bufio"
	"fmt"
	"os"
)

/*

点 i 满足条件 ⟺ 不存在点 j 使得 X_j < X_i 且 Y_j < Y_i,也就是经典的「左下无点 / 帕累托前沿 (staircase)」。

O(N) 实现:因为 X 是排列,用 `pos[x]` 把点按 x 直接排序好,从左到右扫描,维护「至今最小的 y」。当前点的 y 比这个最小值还小时,说明左下方没有点 → 计数,并更新最小值。
- 排列的特性：长度固定，取值范围是 1-N，无重复，无缺失

时间复杂度 O(N)
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)
	count := 0
	pos := make([]int, N+1)
	minY := N + 1

	for i := 1; i <= N; i++ {
		var X, Y int
		fmt.Fscan(reader, &X, &Y)
		pos[X] = Y
	}

	for i := 1; i <= N; i++ {
		y := pos[i]
		if y < minY {
			count++
			minY = y
		}
	}

	fmt.Fprintln(writer, count)

}
