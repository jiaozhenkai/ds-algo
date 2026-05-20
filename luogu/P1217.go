package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	for i := a; i <= b; i++ {
		if i%2 == 0 {
			continue
		}
		if isPalindromes2(i) && isPrime(i) {
			fmt.Println(i)
		}
	}
}
