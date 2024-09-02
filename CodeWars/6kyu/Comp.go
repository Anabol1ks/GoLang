package main

import (
	"fmt"
	"sort"
)

func Comp(array1 []int, array2 []int) bool {
	sort.Ints(array1)
	sort.Ints(array2)
	if len(array1) == len(array2) {
		for i := 0; i < len(array1); i++ {
			if array1[i]*array1[i] != array2[i] {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

func main() {
	var a1 = []int{121, 144, 19, 161, 19, 144, 19, 11}
	var a2 = []int{11 * 11, 121 * 121, 144 * 144, 19 * 19, 161 * 161, 19 * 19, 144 * 144, 19 * 19}
	fmt.Println(Comp(a1, a2))
}
