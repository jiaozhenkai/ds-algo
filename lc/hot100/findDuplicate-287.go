package main

import "sort"

/*
思路：二分
有 n+1 个数，范围在 [1,n]，为了把 n+1 长度的数组填充满，那么数组中必然要有一个重复的数字（类似抽屉原理）。
这就相当于把 n+1 个数映射到[1,n]（注意是 1～n，不是数组长度）肯定会有重复。
对 [1,n] 做二分查找，n 是数组中最大的那个数，这是题目保证的，1<=nums[i]<=n
- 对于 [1,n]的中点 mid，定一个 count[mid] 为数组中 ≤ mid 的元素个数
- 如果有重复在[1,mid] 之内，count[mid] > mid
- 如果没有重复在[1,mid] 之内，那么刚好应该 count[mid] = mid
- 如果有重复在 [mid+1,n] 之内，count[mid] < mid
*/

func findDuplicate1(nums []int) int {
	ans := -1
	n := len(nums)
	l, r := 1, n-1 // n-1 代表取值范围上界
	for l <= r {
		mid := (l + r) >> 1
		count := 0
		for i := 0; i < n; i++ {
			if nums[i] <= mid {
				count++
			}
		}
		if count <= mid {
			l = mid + 1
		} else {
			r = mid - 1
			ans = mid
		}
	}

	return ans
}

/*
思路：两层 for
*/

/*
思路：排序
*/

func findDuplicate(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	ans := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			ans = nums[i]
		}
	}
	return ans
}
