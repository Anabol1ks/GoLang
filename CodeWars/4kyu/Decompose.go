package main

import (
	"fmt"
)

func Decompose(n int64) []int64 {
	result, _ := decomposeHelper(n, n*n)
	return result
}

func decomposeHelper(n int64, remainder int64) ([]int64, bool) {
	if remainder == 0 {
		return []int64{}, true
	}

	for i := n - 1; i > 0; i-- {
		square := i * i
		if square <= remainder {
			if result, found := decomposeHelper(i, remainder-square); found {
				return append(result, i), true
			}
		}
	}

	return nil, false
}

func main() {
	fmt.Println(Decompose(11))
}
