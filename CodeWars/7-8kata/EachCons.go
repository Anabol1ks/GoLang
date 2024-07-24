package main

import "fmt"

func EachCons(arr []int, n int) (res [][]int) {
	for i := 0; i <= len(arr)-n; i++ {
		sr := []int{}
		for j := i; j < i+n; j++ {
			sr = append(sr, arr[j])
		}
		res = append(res, sr)
	}
	return res
}

func main() {
	fmt.Println(EachCons([]int{1, 2, 3, 4, 5, 6, 7}, 3))
}
