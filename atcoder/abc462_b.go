package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	fmt.Scan(&N)

	recv := make([][]int, N+1)

	for i := 1; i <= N; i++ {
		var n int
		fmt.Fscan(reader, &n)
		for j := 0; j < n; j++ {
			var A int
			fmt.Fscan(reader, &A)
			recv[A] = append(recv[A], i)
		}
	}

	for i := 1; i <= N; i++ {
		fmt.Fprintf(writer, "%d ", len(recv[i]))
		for j := 0; j < len(recv[i]); j++ {
			fmt.Fprintf(writer, "%d ", recv[i][j])
		}
		fmt.Fprintln(writer)
	}
}
