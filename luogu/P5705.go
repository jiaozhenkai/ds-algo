package main

import "fmt"

func main() {
	f := int(ReadFloat64() * 10)
	fNum := float64(f % 10)
	f = f / 10

	// res 只保存原来小数位的结果
	res := 0.0
	tmp := 0.1
	for f > 0 {
		fnum := float64(f % 10)
		res += fnum * tmp
		f = f / 10
		tmp = tmp / 10
	}

	fmt.Printf("%g", fNum+res)
}
