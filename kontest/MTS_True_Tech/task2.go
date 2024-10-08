package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	target := []rune{'M', 'T', 'S'}
	var str string
	fmt.Fscan(in, &str)

	idx := 0
	for _, s := range str {
		if s == target[idx] {
			idx++
		}
		if idx == len(target) {
			fmt.Fprintln(out, 1)
			return
		}
	}
	fmt.Fprintln(out, 0)
}
