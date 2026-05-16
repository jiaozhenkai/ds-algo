package main

import "fmt"

func main() {
	m, t, s := ReadInt(), ReadInt(), ReadInt()
	if t == 0 {
		fmt.Println(0)
		return
	}

	if s%t == 0 {
		fmt.Println(max(m-s/t, 0))
	} else {
		fmt.Println(max(m-s/t-1, 0))
	}
}
