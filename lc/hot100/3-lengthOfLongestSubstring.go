package main

import "fmt"

/*
双指针+hash

left 代表某个不重复字串的起始位置，right 代表某个不重复字串的结束位置，right-left+1就是这个不重复字串的长度
right 指针每次右移一位，判断 s[right] 是否在当前的字串中有重复字符

  - 如果没有重复字符，right 继续右移，left 不变
    -- 此时 maxLen=right-left+1 或者 maxLen+=1 来维护当前子串的最大长度

  - 如果有重复字符，这个重复字符有两种可能的位置：
    -- hash[ch] < left
    -- hash[ch] >= left
    -- 对于 hash[ch] < left：代表重复字符没有落在 [left,right] 中，left 无需更新
    -- 对于 hash[ch] >= left：代表重复字符落在了 [left, right] 中，left 需要更新位置，left=hash[ch]+1
    -- 有重复字符时无需更新 maxLen (不写为啥了)
*/
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	maxLen := 0
	left := 0
	hash := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		ch := s[right]
		if idx, ok := hash[ch]; ok && idx >= left {
			left = idx + 1
		} else {
			maxLen = max(maxLen, right-left+1)
		}
		hash[ch] = right
	}
	return maxLen
}

// 以前提交的代码，现在一时难以理解当初为啥整出这么个操作
func lengthOfLongestSubstring1(s string) int {
	res := 0
	var last [128]int
	for i := 0; i < len(last); i++ {
		last[i] = -1
	}

	// m := make(map[byte]int)
	n := len(s)
	start := 0

	for i := 0; i < n; i++ {
		ch := s[i]
		start = max(start, last[ch]+1)
		res = max(res, i-start+1)

		last[ch] = i
	}

	return res
}

func main() {
	s := "tmmzuxt"
	fmt.Println(lengthOfLongestSubstring(s))
	//s1 := "abba"
	//fmt.Println(lengthOfLongestSubstring1(s1))
}
