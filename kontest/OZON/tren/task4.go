package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var l int
		fmt.Fscan(in, &l)
		arr := make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Fscan(in, &arr[j])
		}
		var maxLen int
		var maxSubarray []int

		for i := 0; i < len(arr); i++ {
			counts := make(map[int]int)
			start := i

			for j := i; j < len(arr); j++ {
				counts[arr[j]]++

				if len(counts) > 2 {
					break
				}

				if j-start+1 > maxLen {
					maxLen = j - start + 1
					maxSubarray = arr[start : j+1]
				}
			}
		}
		fmt.Fprintln(out, len(maxSubarray))
	}
}
