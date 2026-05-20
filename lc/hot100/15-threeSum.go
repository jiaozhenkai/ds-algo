package main

import (
	"fmt"
	"sort"
)

/*
思路：排序+双指针；**先对 nums 升序排序**，然后**每一轮循环** L 指针设置为 i+1, R 指针设置为 len(nums)-1;(符合需求的序列中，一定要既有正数也有负数或者就是全部为 0，不可能全部为正数或负数. 假设 sum=nums[i]+nums[L]+nums[R]
- if nums[0] >0 || nums[len(nums)-1] <0 直接 return，都不符合需求否则开启循环遍历。
- if nums[i]> 0， 则直接跳过，因为 nums[i] 以及 nums[i] 之后不可能有3个数字相加等于 0 的情况。
	- 只有当 nums[i] <=0 时才开启循环判断。
- if nums[i] == nums[i-1] 两个数字重复，直接跳过。因为 sum==0 符合要求的，nums[i-1] 也早已经判断过了。
- if sum>0, R--; 这意味着 num[R] 贡献的数字太大，应该往左移动。
- if sum<0, L++; 这意味着 num[L] 贡献的数字太小，应该往右移动。
- if sum==0：
	- if nums[L]==nums[L+1],L++; 因为相同的数字会导致结果重复
	- if nums[R]==nums[R-1],R--; 因为相同的数字会导致结果重复
	- R--,L++; 原因：比如 L,L+1,L+2 分别为 1，1，2；第一次判断之后，L++ 之后 L 会落到第二个 1 上面，再一次判断之后 1和2不相等，但是 L 还是在第二个 1 上，此时第一个和第二个 1 还是有重复，应该再递增一次。
*/

func threeSum1(nums []int) [][]int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	if nums[0] > 0 || nums[len(nums)-1] < 0 || len(nums) < 3 {
		return [][]int{}
	}
	if nums[0] == 0 && nums[len(nums)-1] == 0 {
		return [][]int{{0, 0, 0}}
	}

	ans := [][]int{}

	for i := 0; i < len(nums); i++ {
		// 序列中某一个数字大于 0 和一定大于 0 ，这个和 nums[0] >0 边界情况的判断不一样。
		if nums[i] > 0 {
			break
		}
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 注意不能放在 for 的外面
		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[L], nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1] {
					R--
				}
				L++
				R--
			}
			if sum > 0 {
				R--
			}
			if sum < 0 {
				L++
			}
		}

	}

	return ans

}

// 下面这个是错的，感觉应该可以转换为两数之和。
func threeSum(nums []int) [][]int {
	tmpAns := [][]int{}
	// ans := [][]int{}
	for i, v := range nums {
		threeSumSlice := []int{}
		target := 0 - v
		// two sum problem
		func() {
			hash := make(map[int]int)
			for j := 0; j < len(nums); j++ {
				if j == i {
					continue
				}
				if v, ok := hash[target-nums[j]]; ok {
					threeSumSlice = append(threeSumSlice, target-nums[j])
					threeSumSlice = append(threeSumSlice, v)
				} else {
					hash[nums[j]] = j
				}
			}
			return
		}()

		tmpAns = append(tmpAns, threeSumSlice)
	}

	for _, v := range tmpAns {
		sort.Slice(v, func(i, j int) bool {
			return v[i] < v[j]
		})
	}

	return tmpAns

}
func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}

	fmt.Println(threeSum1(nums))
}
