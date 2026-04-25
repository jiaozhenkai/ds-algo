package main

import "fmt"

/*
前缀和：求出前缀和之后，问题转换为 sum[j]-sum[i]=k (0<=i<j<len(nums)), for 循环遍历前缀和中满足 sum[j]-sum[i]=k 等式的个数，但是这样的时间复杂度是n^2。
之前 twosum 问题中，将问题转换为了 x=target-nums[i]，x 是否在hash表中存在的情况。这里也可以借鉴将问题转换为 sum[j]-k=sum[i];
即遍历到 j 时，[0,j-1] 中有多少个满足 sum[j]-k 的前缀和，同样用 hash 表保存 key:sum[j], value:前缀和为 sum[j] 的个数
- if exist ans++
- if !exist hash[sum[j]-k]++
*/

func subarraySum1(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

	ans := 0
	sum := make([]int, len(nums))
	sum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		sum[i] = sum[i-1] + nums[i]
	}

	hash := make(map[int]int)
	hash[0] = 1 // 如果 sum[i]==k（sum[i]-k==0）这意味着前 i 项和已经满足了需求，因此 hash[0]=1
	for _, s := range sum {
		if cnt, ok := hash[s-k]; ok {
			ans += cnt
		}
		hash[s]++
	}

	return ans
}

/*
前缀和
subarraySum1 里面先求了前缀和，但是根据 twosum 的思路，完全可以一遍遍历，一遍计算前缀和。
*/
func subarraySum2(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}
	ans := 0
	hash := make(map[int]int)
	hash[0] = 1
	sum := 0
	for _, v := range nums {
		sum += v
		if cnt, ok := hash[sum-k]; ok {
			ans += cnt
		}
		hash[sum]++
	}

	return ans
}

/*
暴力思路：两次遍历，要注意：
1. nums[i]==k，是有可能出现的，所以需要计入总数
2. nums[i]==k, 后续num[i+1]==0, 是有可能出现的， 需要继续在内层循环遍历一次。
*/
func subarraySum(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		tmp := nums[i]
		//if tmp == k {
		//	ans++
		//}
		for j := i + 1; j < len(nums); j++ {
			tmp += nums[j]
			if tmp == k {
				ans++
			}
		}
	}

	return ans

}

func main() {
	nums := []int{1, 2, 3}
	a := subarraySum1(nums, 3)
	fmt.Println(a)
}
