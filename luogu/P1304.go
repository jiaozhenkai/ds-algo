package main

import (
	"fmt"
)

//func isPrime(n int) bool {
//	if n < 2 {
//		return false
//	}
//	if n < 4 {
//		return true
//	}
//	if n%2 == 0 || n%3 == 0 {
//		return false
//	}
//	for i := 5; i*i <= n; i += 6 {
//		if n%i == 0 || n%(i+2) == 0 {
//			return false
//		}
//	}
//	return true
//}

func solve(n int) {
	for i := 4; i <= n; i += 2 {
		for j := 2; j <= i/2; j++ {
			if isPrime(i-j) && isPrime(j) {
				fmt.Printf("%d=%d+%d\n", i, j, i-j)
				break
			}
		}

	}
}
func main() {
	var n int
	fmt.Scan(&n)
	solve(n)
}
