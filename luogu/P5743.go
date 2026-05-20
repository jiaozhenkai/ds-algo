package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	rev := 1

	/*
		0  1       2         3         4
		s0  s0/2-1   s1/2-1    s2/2-1     1
		1 时已经吧 s0 求出来了，不用再去迭代一次了因此 i<n-1

	*/
	for i := 0; i < n-1; i++ {
		rev = (rev + 1) * 2
	}
	fmt.Println(rev)
}
