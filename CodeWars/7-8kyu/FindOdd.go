package main

import (
	"fmt"
)

func FindOdd(seq []int) int {
	res := 0
	for _, x := range seq {
		res ^= x
	}
	return res
}

func main() {
	fmt.Println(FindOdd([]int{20, 1, -1, 2, -2, 3, 3, 5, 5, 1, 2, 4, 20, 4, -1, -2, 5}))
}
