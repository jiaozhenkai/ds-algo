package main

/*
要想储存最大水量，就是要让横纵坐标都是最大；所以乘积会最大。
- 数组的位置代表了横坐标，元素的值代表了纵坐标。

思路：双指针
left 指针处所在的元素是 x, right 指针处所在的元素是 y; x<y（简单期间可以假设 y=x+1）; 此时面积为 (right-left)*x。

接下来重点是考虑怎么移动 left 和 right 指针，直觉上移动 left 指针是对的，定量定性分析：
1.假设现在要把 right 左移一位，对应的元素是 y1
- if y1 < x, 面积一定会变小，因为长和高都变小了。
- if y1 > x, 面积也一定会变小，因为长变为了 right-left-1，变小了。
- 无论 y1 是否大于 y 都无法改变之前分析的两个结果。
- 所以左移 right 一定不行: 只要左移 right 指针，一定会导致面积变小。

2. 考虑 left 右移一位，对应的元素是 x1 且 x1>x （简单期间可以假设 x1=x+1）
- if x1 < y, 面积变为 (right-left-1)*x1 >= (right-left)*x
- if x1 >=y, 面积变为（right-left-1)*y >= (right-left)*x
3. 考虑 left 右移一位，对应的元素是 x1 且 x1<=x （简单期间可以假设 x1=x-1）
- 无论如何，面积一定会变小

上面 3 步分析得出：只有 left 右移，才可能获得更大的面积。


思路：暴力算法
两层 for 循环
*/

// 暴力算法,会超出时间限制
func maxArea(height []int) int {
	maxarea := 0

	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			// 注意这里不是 j-i+1, 求面积要看 j和i之间的空位而不是数字个数。
			area := (j - i) * min(height[i], height[j])
			maxarea = max(maxarea, area)
		}
	}
	return maxarea
}

// 双指针
func maxArea1(height []int) int {
	maxA := 0
	left := 0
	for right := len(height) - 1; left < right; {
		maxA = max(maxA, (right-left)*min(height[left], height[right]))
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return maxA
}
