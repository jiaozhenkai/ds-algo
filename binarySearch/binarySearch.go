package main

import "fmt"

// 递归
func bs(l, r, target int, arr []int) int {
	if l > r {
		return -1
	}
	// 偶数的中间放到左边
	mid := (r-l)/2 + l
	if arr[mid] == target {
		return mid
	} else if arr[mid] > target {
		r = mid - 1
		return bs(l, r, target, arr)
	} else {
		l = mid + 1
		return bs(l, r, target, arr)
	}
}

// 迭代
func bs1(l, r, target int, arr []int) int {
	for l <= r {
		mid := (r-l)/2 + l
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

// 找第一个 >T 的位置
func bs2(arr []int, target int) int {
	l, r := 0, len(arr) // 左闭右开
	for l < r {
		mid := (r-l)/2 + l
		if arr[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	i := bs(0, len(arr)-1, 7, arr)
	fmt.Println(i, arr[i])

	ii := bs1(0, len(arr)-1, 7, arr)
	fmt.Println(ii, arr[ii])

	j := bs2(arr, 4)
	fmt.Println(j)
}
