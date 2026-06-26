package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N)

	H, L := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Fscan(reader, &H[i], &L[i])
	}

	suffixMax := make([]int, N)
	suffixMax[N-1] = H[N-1]
	for i := N - 2; i >= 0; i-- {
		if H[i] > suffixMax[i+1] {
			suffixMax[i] = H[i]
		} else {
			suffixMax[i] = suffixMax[i+1]
		}
	}

	var Q int
	fmt.Fscan(reader, &Q)

	//bs := func(T int) int {
	//	l, r := 0, N
	//	for l < r {
	//		mid := (r-l)/2 + l
	//		if L[mid] > T {
	//			r = mid
	//		} else {
	//			l = mid + 1
	//		}
	//	}
	//	return l
	//}

	for i := 0; i < Q; i++ {
		var T int
		fmt.Fscan(reader, &T)
		// pos := bs(T)
		pos := sort.SearchInts(L, T+1)
		fmt.Fprintf(writer, "%d\n", suffixMax[pos])
	}

}
