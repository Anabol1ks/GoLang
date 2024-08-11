package main

import (
	"bufio"
	"fmt"
	"os"
)

const MOD = 1000000007

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Fscan(in, &n)
		arr_1 := make([]int, n)
		arr_2 := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &arr_1[j])
		}
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &arr_2[j])
		}
		res := 1
		for j := 0; j < n; j++ {
			l := arr_1[j]
			r := arr_2[j]
			i := j + 1
			min := (l + i - 1) / i
			max := r / i
			if min > max {
				res = 0
				break
			}
			res = res * (max - min + 1) % MOD
		}
		fmt.Fprintln(out, res)
	}
}
