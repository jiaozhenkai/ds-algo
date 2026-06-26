package main

/*
思路：hash，以序列中的某个元素 x 未开始，找是否存在 x+1,x+2 的数字，是否存在就可以用 hash 来快速判断。

- 但是以 x 开始的元素，其序列可能并不是最长的，因为可能还存在 x-1,所以：
  - 如果 x-1 存在，那么直接跳过等遇到 x-1 时再去判断。
	- 这样做也不会有遗漏，比如 1，1，1，2，3，4，如果 for hash 表先遇到了 1 那么就可以找到最长的序列，如果先遇到了 2，3，4
		那么最终到 1 时也可以正确判断。
  - 如果 x-1 不存在，直接开始从 x 判断找 x+1，x+2 是否存在即可。

- 需要遍历 hash 表，而不是遍历 nums 切片，比如  1，1，1，2，3，4，如果遍历的话会 1 会遍历 3 次，做了很多无用的工作。
*/

func longestConsecutive(nums []int) int {
	res := 0
	hash := make(map[int]struct{})
	for _, v := range nums {
		hash[v] = struct{}{}
	}

	for k, _ := range hash {
		count := 1
		if _, ok := hash[k-1]; ok {
			continue
		}
		end := k + 1
		for {
			if _, ok := hash[end]; ok {
				end++
				count++
			} else {
				break
			}
		}
		res = max(res, count)
	}

	return res
}
