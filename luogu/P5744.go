package main

import "fmt"

type Stu struct {
	name  string
	age   int
	score int
}

func main() {
	var n int
	fmt.Scan(&n)
	stus := []Stu{}
	for i := 0; i < n; i++ {
		var name string
		var age int
		var score int
		var stu Stu
		fmt.Scan(&name, &age, &score)
		stu.name = name
		stu.age = age
		stu.score = score
		stus = append(stus, stu)
	}

	for _, v := range stus {
		score := int(float64(v.score) * 1.2)
		if score > 600 {
			score = 600
		}
		fmt.Println(v.name, v.age+1, score)
	}

}
