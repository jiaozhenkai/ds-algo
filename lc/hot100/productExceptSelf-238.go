package main

import "fmt"

/*
思路：既然返回数组不计入空间复杂度，那么可以考虑让 ans 作为左右乘积数组之一，借助返回数组 ans，优化空间复杂度为O（1）.
ans[i] 记录 nums[i] 的左乘积列表。左乘积列表已经知道，只需要想办法获取右乘积列表或者直接计算出来结果即可。
ans[i] 如果要表示为最终结果，那么 ans[i] = rp(i+1)*nums[i+1],考虑：
a,b,c,d,e
- e 的右乘积列表 rp(e) 为 1，ans[e](最终答案) = ans[e](e 的左乘积列表)*rp(e)
- d 的右乘积列表 rp(d) 为 e*rp(e)，ans[d](最终答案) = ans[d](d 的左乘积列表)*rp(d)
- c 的右乘积列表 rp(c) 为 d*rp(d)，ans[c](最终答案) = ans[c](c 的左乘积列表)*rp(c)
- ...
- 以上可以看出，只需要有一个临时变量保存 rp(e),rp(d),rp(c) 即可。而 rp(e),rp(d) 等都可以直接算出来。

这个方法不够清晰。。。。
*/
func productExceptSelf1(nums []int) []int {
	// answer, 也是 lp 左乘积列表。
	ans := make([]int, len(nums))
	ans[0] = 1
	for i := 1; i < len(nums); i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	// 右乘积列表；nums[len(nums)-1]的右乘积列表初始化为1
	rp := 1

	for i := len(nums) - 1; i >= 0; i-- {
		ans[i] = ans[i] * rp
		rp *= nums[i]
	}
	return ans
}

/*
思路：左右乘积列表
a,b,c,d,e
- a 左边没有数字，那 a 的左乘积列表为 1； 同理 e 的右乘积列表为 1
- c 的左乘积列表 = b*b的左乘积列表。(这就就是前缀和的思想)
*/
func productExceptSelf(nums []int) []int {
	ans := make([]int, len(nums))
	// lp[i] 代表 nums[i] 的左乘积列表
	lp := make([]int, len(nums))
	// rp[i] 代表 nums[i] 的右乘积列表
	rp := make([]int, len(nums))

	lp[0] = 1
	//rp[len(nums)-1] = nums[len(nums)-1]
	for i := 1; i < len(nums); i++ {
		lp[i] = lp[i-1] * nums[i-1]
	}

	rp[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		rp[i] = rp[i+1] * nums[i+1]
	}

	for i := 0; i < len(nums); i++ {
		ans[i] = lp[i] * rp[i]
	}
	fmt.Println(ans, lp, rp)
	return ans
}

func main() {
	nums := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(nums))
}
