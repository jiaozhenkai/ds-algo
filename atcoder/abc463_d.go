package main

import (
	"bufio"
	"os"
)

/*
本质：最大化最小值，最小值由 ｜p-q｜ 给出。
*/
func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

}
