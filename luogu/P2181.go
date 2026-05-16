package main

import "fmt"

func main() {
	var n uint64
	fmt.Scan(&n)
	// 注意溢出
	res := n * (n - 1) / 2 * (n - 2) / 3 * (n - 3) / 4
	fmt.Println(res)
}
