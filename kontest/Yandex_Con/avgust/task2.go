package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func med(arr []int) int {
	n := len(arr)
	sorted := make([]int, n)
	copy(sorted, arr)
	sort.Ints(sorted)
	return sorted[n/2]
}

func countBmed(arr []int, B int) int {
	n := len(arr)
	count := 0

	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if (j-i+1)%2 == 1 {
				subarray := arr[i : j+1]
				if med(subarray) == B {
					count++
				}
			}
		}
	}
	return count
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, b int
	fmt.Fscan(in, &n, &b)
	arr := make([]int, n)
	for i := range arr {
		var num int
		fmt.Fscan(in, &num)
		arr[i] = num
	}
	result := countBmed(arr, b)
	fmt.Fprintln(out, result)
}
