package main

import (
	"fmt"
	"slices"
)

/*
[]int{1, 2, 3, 4, 5, 6, 7}, k=3

	|
	V

[]int{5, 6, 7, 1, 2, 3, 4}
*/
func rotate3(nums []int, k int) {
	if len(nums) == 1 || len(nums) == k {
		return
	}

	if k > len(nums) {
		k = k % len(nums)
	}

	slices.Reverse(nums)
	slices.Reverse(nums[0:k])
	slices.Reverse(nums[k:len(nums)])
}

/*
思路：三次反转
序列 Ai...Ak-1,Ak,Ak+1...Aj, 最终目的是得到 Ak,Ak+1...Aj,Ai...Ak-1
先整体翻转得到 Aj...Ak+1,Ak,Ak-1...Ai
- 前半部分总共 k 个数字，后半部分总共 len(nums）-k 个
再翻转前部分得到 Ak,Ak+1...Aj
再翻转后部分得到 Ai...Ak-1
*/
func rotate2(nums []int, k int) {
	if len(nums) == 1 || k == len(nums) {
		return
	}

	if k > len(nums) {
		k = k % len(nums)
	}

	reverse := func(i, j int) {
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
	}

	reverse(0, len(nums)-1)
	reverse(0, k-1)
	reverse(k, len(nums)-1)

}

/*
还是和 rotate 一样的思路，但是写法更加简洁。比如 nums := []int{1, 2, 3, 4, 5, 6, 7}, k=3，长度 n=7
nums 中的数字 ｜ nums 中的下标 ｜ 旋转之后新数组中的数字对应的旧数组中的的数字的下标
1｜0｜4
2｜1｜5
3｜2｜6
4｜3｜0
5｜4｜1
6｜5｜2
7｜6｜3
比如 nums[0] 要复制给第三列0所在的位置，也就是newNums[3]
可以看到第三列，就是做了一个旋转，而旋转可以用取模来表示，这里是下标最大到 6 就归0，因此可以模 7 ；由于周期性的存在， 8%7 和 1%7 没有什么区别，所以被除数可以一直递增。可以从上面第三列看到规律：(i+k)%7 (i是未经过旋转的数组的下标)符合要求。
*/
func rotate1(nums []int, k int) {
	if len(nums) == k {
		return
	}
	newNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		newNums[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, newNums)
}

/*
思路：
k == len(nums) 意味着经过了一个周期，数组不会变直接返回即可。只有在 k<len(nums)时才需要有转动。如果 k>len(nums) 那么经过了多轮周期，相当于k=k%len(nums) 次旋转

空间复杂度是 O(n)
*/
func rotate(nums []int, k int) {
	if len(nums) == 1 || k == len(nums) {
		return
	}

	if k > len(nums) {
		k = k % len(nums)
	}
	tmp := make([]int, k)
	for i := 0; i < k; i++ {
		tmp[i] = nums[i+len(nums)-k]
	}

	idx := len(nums) - 1
	for i := len(nums) - k - 1; i >= 0; i-- {
		nums[idx] = nums[i]
		idx--
	}
	for i := 0; i < k; i++ {
		nums[i] = tmp[i]
	}
}
func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	rotate3(nums, 3)
	fmt.Println(nums)
	nums1 := []int{1, 2}
	rotate3(nums1, 7)
	fmt.Println(nums1)
}
