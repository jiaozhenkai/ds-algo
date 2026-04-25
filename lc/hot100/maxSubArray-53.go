package main

/*
思路：动态规划
假设 i<j, 以num[i]结尾的子数组的最大和不受之后 num[j] 的计算，无后效性。
假设 dp[i] 是以 nums[i] 结尾的子数组的最大和（如果nums中的数字都是>=0的，那么 dp[i]>=dp[i-1]+nums[i] 一定是恒成立的）；由于题目中数字有正有负，此时 dp[i-1] 可能有几种状态：
- dp[i-1] <=0；此时 dp[i-1] + nums[i] 只能使得和更小或者不变。
- dp[i-1] >0;此时 dp[i-1]+nums[i] 一定会让和更大。
因此状态转移方程为：dp[i] = max(dp[i-1]+nums[i], nums[i])
*/
func maxSubArray1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] <= 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
	}
	ans := dp[0]
	for i := 1; i < len(dp); i++ {
		ans = max(ans, dp[i])
	}
	return ans
}

// 动态规划，空间复杂度O(1)；maxSubArray1 中，其实只需要记录某一轮循环中最大的那个值就行。
func maxSubArray2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// preMax 标识 nums[i] 轮时的连续子数组的最大值，要么是 preMax+nums[i],要么是nums[i]
	preMax := nums[0]
	// ans 表示当前连续子数组的最大值，要么是已经存在的 ans, 要么是现在计算出来的 preMax
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		preMax = max(preMax+nums[i], nums[i])
		ans = max(ans, preMax)
	}
	return ans
}

// 分治
func maxSubArray3(nums []int) int {
	return 0
}

// 暴力解法，会超时。
func maxSubArray(nums []int) int {
	ans := -0xFFFFFFFF
	for i := 0; i < len(nums); i++ {
		tmp := 0
		for j := i; j < len(nums); j++ {
			tmp += nums[j]
			ans = max(ans, tmp)
		}
	}

	return ans
}
