// 前后缀和/最值

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	n := len(arr)

	// 前缀和
	prefix := make([]int, n)
	prefix[0] = arr[0]
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1] + arr[i]
	}
	fmt.Println(prefix)

	// 后缀和
	suffix := make([]int, n)
	suffix[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		suffix[i] = suffix[i+1] + arr[i]
	}
	fmt.Println(suffix)

	// 前缀最大值
	prefixMax := make([]int, n)
	prefixMax[0] = arr[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = max(prefixMax[i-1], arr[i])
	}
	fmt.Println(prefixMax)

	// 后缀最大值
	suffixMax := make([]int, n)
	suffixMax[n-1] = arr[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMax[i] = max(suffixMax[i+1], arr[i])
	}
	fmt.Println(suffixMax)

	// 前缀区间和 [l,r]
	l, r := 2, 5
	sum := 0
	if l == 0 {
		sum = prefix[r]
	} else {
		sum = prefix[r] - prefix[l-1]
	}
	fmt.Println(sum)

	// 后缀区间和 [l,r]
	sum1 := 0
	if r == n-1 {
		sum1 = suffix[l]
	} else {
		sum1 = suffix[l] - suffix[r+1]
	}
	fmt.Println(sum1)
}
