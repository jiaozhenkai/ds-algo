package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

/*
没有遇到 E 就一直持续统计，遇到了 E 就统计上次结束到 E 时的比分
11 分制和 21 分制都是一样的。
*/

// 一方先到 11/21 分，且分差大于 2 分才停止。
func isGameOver(w, l, target int) bool {
	if (w >= target || l >= target) && math.Abs(float64(w-l)) >= 2 {
		return true
	}
	return false
}
func main() {
	count11W, count11L := 0, 0
	count21W, count21L := 0, 0
	scanner := bufio.NewScanner(os.Stdin)
	output21 := []string{}
	done := false
	for scanner.Scan() {
		if done { // 要 break 掉，不然还需要再来一次回车
			break
		}
		line := scanner.Text()
		for _, v := range line {
			if v == 'E' {
				fmt.Printf("%d:%d\n\n", count11W, count11L)
				output21 = append(output21, fmt.Sprintf("%d:%d", count21W, count21L))
				done = true
				break
			}
			if v == 'W' {
				count11W++
				count21W++
			}
			if v == 'L' {
				count11L++
				count21L++
			}
			// 11
			if isGameOver(count11W, count11L, 11) {
				fmt.Printf("%d:%d\n", count11W, count11L)
				count11L = 0
				count11W = 0
			}
			//21
			if isGameOver(count21W, count21L, 21) {
				output21 = append(output21, fmt.Sprintf("%d:%d", count21W, count21L))
				count21L = 0
				count21W = 0
			}
		}
	}
	for _, v := range output21 {
		fmt.Println(v)
	}
}
