package main

import (
	"fmt"
	"sort"
)

/*
定长滑动窗口思路：如果 s 的字串 t 中每个字符出现的数量和 p 中每个字符的数量一样，则符合需求。
因为都是小写字母，最多 26 长度的数组就可以全部存储下来。而数组可以直接比较是否相等。
*/
func findAnagrams1(s string, p string) []int {
	ans := []int{}
	cntP := [26]int{}
	cntS := [26]int{}
	for _, v := range p {
		cntP[v-'a'] += 1
	}
	lenP := len(p)
	L := 0
	R := 0
	for R < len(s) {
		cntS[s[R]-'a']++
		if R-L+1 < lenP { // 窗口大小不够。
			R++
			continue
		}
		if cntP == cntS {
			ans = append(ans, L)
		}
		cntS[s[L]-'a']-- //注意不能直接设置为0，因为窗口中还有其他的相同字符。只需要减1保证最左窗口的字符离开窗口即可。
		L++
		R++
	}

	return ans
}

/*
暴力思路：滑动窗口+排序。窗口大小为 len(p)
先对 p 按照字符排序，再在每次的循环窗口中对 len(p) 长度的 s 的字串进行排序。然后对比排序后的子串和p进行对比，但是这样时间复杂度为On^2(logn)
*/

// 会超出时间限制。
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}
	ans := []int{}
	L := 0
	R := len(p) - 1

	byteP := []byte(p)
	sort.Slice(byteP, func(i, j int) bool {
		return byteP[i] < byteP[j]
	})

	for R < len(s) {

		byteS := []byte(s[L : R+1])

		sort.Slice(byteS, func(i, j int) bool {
			return byteS[i] < byteS[j]
		})
		fmt.Println(string(byteS), string(byteP))
		if string(byteP) == string(byteS) {
			ans = append(ans, L)
		}

		L++
		R++
	}
	return ans
}

func main() {
	//s := "eklpyqrbgjdwtcaxzsnifvhmoueklpyqrbgjdwtcaxzsnifvhmou"
	//p := "yqrbgjdwtcaxzsnifvhmou"
	s := "cbaebabacd"
	p := "abc"
	fmt.Println(findAnagrams1(s, p))
}
