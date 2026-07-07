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

	var S string
	fmt.Fscan(reader, &S)
	bs := []byte(S)
	o := []byte{}
	for i := 0; i < len(bs); i++ {
		if bs[i] <= '9' && bs[i] >= '0' {
			o = append(o, bs[i])
		}
	}

	fmt.Fprintln(writer, string(o))
}
