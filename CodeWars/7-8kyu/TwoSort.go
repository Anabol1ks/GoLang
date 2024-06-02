package main

import (
	"fmt"
	"sort"
)

func TwoSort(arr []string) (res string) {
	sort.Strings(arr)
	for w, i := range arr[0] {
		res += string(i)
		if w != len(arr[0])-1 {
			res += "***"
		}
	}
	return
}

func main() {
	s := []string{"bitcoin", "take", "over", "the", "world", "maybe", "who", "knows", "perhaps"}
	fmt.Println(TwoSort(s))
}
