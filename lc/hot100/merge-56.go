package main

import (
	"fmt"
	"sort"
)

/*
思路：按照左端点排序，排序之后，merged[0] = intervals[0]
如果 intervals[j][0] <= intervals[i][1] (i<j) 则证明有重复：
- 更新 merged[i][1]=max(intervals[i][1], intervals[j][1])
如果 intervals[j][0] > intervals[i][1] (i<j) 则证明无重复：
- 将 intervals[j] append 到 merged 数组

因为已经按照左端点进行过排序，所以 intervals[j][0] >= intervals[i][0] 是肯定成立的，所以只需要判断右端点。
*/
func merge(intervals [][]int) [][]int {
	merged := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	merged = append(merged, intervals[0])

	for i := 1; i < len(intervals); i++ {
		mergedLen := len(merged)
		if intervals[i][0] <= merged[mergedLen-1][1] { // 重复
			merged[mergedLen-1][1] = max(merged[mergedLen-1][1], intervals[i][1])
		} else {
			merged = append(merged, intervals[i])
		}
	}
	return merged
}
func main() {
	nums := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	nums1 := [][]int{{4, 7}, {1, 4}}
	nums2 := [][]int{{2, 3}, {4, 5}, {6, 7}, {8, 9}, {1, 10}}
	fmt.Println(merge(nums))
	fmt.Println(merge(nums1))

	fmt.Println(merge(nums2))
}
