package main

import (
	"fmt"
)

func ArrayDiff(a, b []int) (res []int) {
	el := make(map[int]bool)
	for _, i := range b {
		el[i] = true
	}
	for _, i := range a {
		if !el[i] {
			res = append(res, i)
		}
	}
	return
}

func main() {
	fmt.Println(ArrayDiff([]int{1, 2, 2, 3, 1, 5}, []int{1, 2}))
}
