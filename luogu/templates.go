package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

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

func ReadStringFromLine() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		fmt.Println(line) // 处理每一行
	}
	/*
		var n, m int
		fmt.Scan(&n, &m)
		scanner := bufio.NewScanner(os.Stdin)
		arr := [102][102]byte{}
		for i := 1; i <= n; i++ {
			scanner.Scan()
			line := scanner.Text()
			for j := 1; j <= m; j++ {
				arr[i][j] = line[j-1]
			}
		}
		fmt.Println(arr)
	*/
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

// 试除法判断是否是素数
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n < 4 {
		return true
	}
	if n%2 == 0 || n%3 == 0 {
		return false
	}
	for i := 5; i <= int(math.Sqrt(float64(n))); i += 6 {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
	}
	return true
}

// 判断一个数字是不是回文数，但是引入了额外空间
func isPalindromes(n int) bool {
	if n < 10 {
		return true
	}
	arr := []int{}
	for n > 0 {
		num := n % 10
		arr = append(arr, num)
		n = n / 10
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		if arr[i] != arr[j] {
			return false
		}
	}
	return true
}

// 判断一个数字是不是回文数
/*
反转之后是不是等于原数字
*/
func isPalindromes1(n int) bool {
	if n < 0 || (n%10 == 0 && n != 0) {
		return false
	}
	original := n
	rev := 0
	// n 代表还没有被处理的高位部分
	// rev 代表已经处理并且反转过后的低位部分
	for n > 0 {
		rev = rev*10 + n%10
		n = n / 10
	}
	return rev == original
}

// 判断一个数字是不是回文数
/*
反转之后是不是等于原数字，但是不用全部反转，反转到一半就行了。
*/
func isPalindromes2(n int) bool {
	if n < 0 || (n%10 == 0 && n != 0) {
		return false
	}
	rev := 0
	// n 代表还没有被处理的高位部分
	// rev 代表已经处理并且反转过后的低位部分
	for n > rev {
		rev = rev*10 + n%10
		n = n / 10
	}
	// 偶数位 || 奇数位
	return n == rev || n == rev/10
}
