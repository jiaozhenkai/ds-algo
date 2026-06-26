// 有问题
package main

import "fmt"

const (
	toInner = iota
	toOutter
)

type pair struct {
	dir    int
	career string
}

func main() {
	var n, m int
	pairs := []pair{}
	fmt.Scan(&n, &m)
	for i := 0; i < n; i++ {
		var dir int
		var c string
		var p pair
		fmt.Scan(&dir, &c)
		p.dir = dir
		p.career = c
		pairs = append(pairs, p)
	}

	/*
			4 种组合
			a = 0 && pair.dir = 0 -> 顺时针
		    a = 0 && pair.dir = 1 -> 逆时针
			a = 1 && pair.dir = 0 -> 逆时针
			a = 1 && pair.dir = 1 -> 顺时针
		    定义逆时针为正方向，也就是数组从左往右
	*/
	pos := 0
	for i := 0; i < m; i++ {
		var a int
		var s int
		fmt.Scan(&a, &s)

		if (a == 0 && pairs[pos].dir == 0) || (a == 1 && pairs[pos].dir == 1) {
			if pos >= s {
				pos = pos - s
			} else {
				tmp := s - pos
				pos = len(pairs) - tmp
			}
		}
		if (a == 0 && pairs[pos].dir == 1) || (a == 1 && pairs[pos].dir == 0) {
			if len(pairs)-1-pos >= s {
				pos += s
			} else {
				tmp := len(pairs) - 1 - pos
				s -= tmp
				pos = s
			}
		}
	}

	fmt.Println(pairs[pos].career)
}
