package main

import "sort"

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		if envelopes[i][0] != envelopes[j][0] {
			return envelopes[i][0] < envelopes[j][0]
		}
		return envelopes[i][1] > envelopes[j][1]
	})
	tails := []int{}
	bs := func(target int) int {
		l, r := 0, len(tails)-1
		for l <= r {
			mid := (r-l)/2 + l
			if tails[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return l
	}

	for _, v := range envelopes {
		pos := bs(v[1])
		if pos == len(tails) {
			tails = append(tails, v[1])
		} else {
			tails[pos] = v[1]
		}
	}
	return len(tails)
}
