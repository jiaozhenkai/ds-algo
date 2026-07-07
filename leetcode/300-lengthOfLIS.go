package main

import "slices"

/*
dp：
dp[i]=1+max{dp[j]∣j<i 且 nums[j]<nums[i]}
*/
func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	for i := 0; i < len(dp); i++ {
		dp[i] = 1
	}

	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	return slices.Max(dp)
}

/*
贪心 +二分
*/

func lengthOfLIS1(nums []int) int {
	tails := []int{}

	bisectLeft := func(target int) int {
		l, r := 0, len(tails)-1
		for l <= r {
			mid := (r-l)/2 + l
			if tails[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}

	for i := 0; i < len(nums); i++ {
		pos := bisectLeft(nums[i])
		if pos == len(tails) {
			tails = append(tails, nums[i])
		} else {
			tails[pos] = nums[i]
		}
	}

	return len(tails)
}
