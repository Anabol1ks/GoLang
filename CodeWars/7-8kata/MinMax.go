package main

import (
	"fmt"
	"sort"
)

func MinMax(arr []int) [2]int {
	sort.Ints(arr)
	return [2]int{arr[0], arr[len(arr)-1]}
}

func main() {
	fmt.Println(MinMax([]int{1, 2, 3, 4, 5}))
}
