package main

import "fmt"

func main() {
	c := ReadChar()
	res := c - 'a' + 'A'
	fmt.Printf("%c", res)
}
