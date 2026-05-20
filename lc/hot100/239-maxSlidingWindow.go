package main

/*
会超出时间限制
*/

func maxSlidingWindow(nums []int, k int) []int {
	if k == 1 {
		return nums
	}

	if k == len(nums) {
		tmp := nums[0]
		for i := 1; i < len(nums); i++ {
			if tmp < nums[i] {
				tmp = nums[i]
			}
		}
		return []int{tmp}
	}

	res := []int{}

	for i := k - 1; i < len(nums); i++ {
		tmp := nums[i]
		for j := i; j > i-k; j-- {
			if tmp < nums[j] {
				tmp = nums[j]
			}
		}
		res = append(res, tmp)
	}

	return res
}
