package main

import (
	"fmt"
	"sort"
)

func MergeArrays(arr1, arr2 []int) (finArr []int) {
	sortArr := []int{}
	sortArr = append(sortArr, arr1...)
	sortArr = append(sortArr, arr2...)
	sort.Ints(sortArr)
	for i, num := range sortArr {
		if i == 0 || num != sortArr[i-1] {
			finArr = append(finArr, num)
		}
	}
	return
}

func main() {
	fmt.Println(MergeArrays([]int{1, 3, 5, 7, 9, 11, 12}, []int{1, 2, 3, 4, 5, 10, 12}))
}
