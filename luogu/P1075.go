package main

import "fmt"

func main() {
	var n uint32
	fmt.Scan(&n)
	var i uint32 = 0
	for i = 2; i < n; i++ {
		if n%i == 0 {
			fmt.Println(n / i)
			break
		}
	}
}
