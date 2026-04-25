package main

/*
思路：hash
i 处的元素要么和 i 之前的元素相加符合要求，要么和 i 之后的元素相加符合要求
- 和之前的元素符合要求，那么只需看 target-nums[i] 元素是否存在即可。
	- 所以直接让 target-nums[i] 作为 key 放到 hash 表里面即可。
- 和之后的元素符合要求，那么转换为 i 之后的元素和他之前的元素相加是否匹配。
*/

func twoSum(nums []int, target int) []int {
	ans := []int{}

	hash := map[int]int{}

	for idx, val := range nums {

		if v, ok := hash[target-val]; ok {
			ans = append(ans, idx)
			ans = append(ans, v)
			break
		} else {

			hash[val] = idx
		}
	}

	return ans
}
