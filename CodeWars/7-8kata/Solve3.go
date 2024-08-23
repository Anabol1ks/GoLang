package main

import (
	"fmt"
	"sort"
)

func Solve(s string) bool {
	if len(s) > 1 {
		arr := []int{}
		for _, ch := range s {
			arr = append(arr, int(ch-96))
		}
		sort.Ints(arr)
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1]-arr[i] != 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(Solve("abce"))
}
