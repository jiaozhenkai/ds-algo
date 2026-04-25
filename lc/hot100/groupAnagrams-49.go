package main

import (
	"fmt"
	"sort"
)

/*
思路：hash
- 字母异位词成立的充要条件是：两个字符串中的字符都能在组层对方字符串中的字符列表里，则符合需求。

- 将给出的列表中的每个字符串按照字符顺序做排序和去重，hash 去重，然后对 bytes 排序，按照字符顺序从 hash 表中取出来字符，组合成新的字符串。

- 如果 hash[removeDup(str)] 存在，则证明 str 已经有字母异位词，加入到对应的列表中即可。
- 如果 hash[removeDup(str)] 不存在，则证明 str 还没有字母异位词，removeDup(str) 作为 key, []string{str} 作为 value 插入 hash 表。
- 最后 for 遍历 hash ，去除 value,写入二维 string 切片将结果返回。

- hash 表的 key 是去重之后的字符串即：removeDup(str)，value 是 []string.
*/

func removeDup(str string) string {
	bytes := []byte(str)
	hash := make(map[byte]struct{})
	for _, v := range bytes {
		if _, ok := hash[v]; ok {
			continue
		} else {
			hash[v] = struct{}{}
		}
	}
	sort.Slice(bytes, func(i, j int) bool {
		return bytes[i] < bytes[j]
	})

	ans := ""
	for _, b := range bytes {
		ans += string(b)
	}
	return ans

}

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	if len(strs) == 0 {
		return ans
	}
	if len(strs) == 1 {
		return [][]string{{strs[0]}}
	}

	hash := make(map[string][]string)
	for _, v := range strs {
		tmpStr := removeDup(v)
		if _, ok := hash[tmpStr]; ok {
			hash[tmpStr] = append(hash[tmpStr], v)
		} else {
			tmp := []string{v}
			hash[tmpStr] = tmp
		}
	}

	for _, v := range hash {
		ans = append(ans, v)
	}
	return ans

}

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagrams(strs))
}
