package main

/*
思路：双指针

left 指针代表当前可以被插入的位置，有两种可能：
- 0 元素的位置。
- 非 0 元素的位置，但是这个非 0 元素是已经被交换过的。

right 指针代表非 0 元素所在的位置

[left, right-1] 这个区间内就是可以被插入的所有位置。

*/

func moveZeroes(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			continue
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}
