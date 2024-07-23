package main

import "fmt"

func Solve(arr []int) []int {
	seen := make(map[int]bool)
	prres := []int{}
	for i := len(arr) - 1; i >= 0; i-- {
		if !seen[arr[i]] {
			seen[arr[i]] = true
			prres = append(prres, arr[i])
		}
	}
	res := []int{}
	for i := len(prres) - 1; i >= 0; i-- {
		res = append(res, prres[i])
	}
	return res
}

func main() {
	fmt.Println(Solve([]int{3, 4, 4, 3, 6, 3}))
}
