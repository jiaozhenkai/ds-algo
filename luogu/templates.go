package main

import "fmt"

func ReadInt() int {
	var n int
	fmt.Scan(&n)
	return n
}

func ReadInt64() int64 {
	var n int64
	fmt.Scan(&n)
	return n
}

// 读有意义的字符（如迷宫的 # .）
func ReadChar() byte {
	var s string
	fmt.Scan(&s)
	return s[0]
}

// 需要精确读取每个字节（含换行符）
func ReadRawChar() byte {
	var c byte
	fmt.Scanf("%c", &c)
	return c
}

func ReadString() string {
	var s string
	fmt.Scan(&s)
	return s
}

func ReadFloat32() float32 {
	var f float32
	fmt.Scan(&f)
	return f
}

func ReadFloat64() float64 {
	var f float64
	fmt.Scan(&f)
	return f
}
