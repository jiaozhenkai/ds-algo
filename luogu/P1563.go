package main

import "fmt"

const (
	toInner = iota
	toOutter
)

type pair struct {
	direction int
	career    string
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
		p.direction = dir
		p.career = c
		pairs = append(pairs, p)
	}
}
