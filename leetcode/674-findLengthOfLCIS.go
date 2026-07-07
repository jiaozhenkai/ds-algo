package main

import (
	"fmt"
	"slices"
)

func findLengthOfLCIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	l := len(nums)
	for i := 1; i < l; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + 1
		}
	}

	return slices.Max(dp)

}

func main() {
	arr := []int{1, 3, 5, 4, 7}
	res := findLengthOfLCIS(arr)
	fmt.Println(res)
}
