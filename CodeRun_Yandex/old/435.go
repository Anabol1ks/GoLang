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

	var n, num, maxCount int
	fmt.Fscan(in, &n)

	for i := 1; i <= n; i++ {
		countDel := 0
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				countDel++
				if j != i/j {
					countDel++
				}
			}
			if countDel >= maxCount {
				maxCount = countDel
				num = i
			}
		}
	}
	fmt.Fprintln(out, num, maxCount)
}
