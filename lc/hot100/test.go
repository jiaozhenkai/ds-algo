package main

import "fmt"

type A struct {
	i int
	j int
}

func main() {
	a := A{i: 0, j: 1}
	hash := map[A]struct{}{}
	hash[a] = struct{}{}

	b := A{
		i: 1,
		j: 0}
	if v, ok := hash[b]; ok {
		fmt.Println(v)
	}

	for k, v := range hash {
		fmt.Println(k, v)
	}

}
