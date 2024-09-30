package main

import (
	"fmt"
	"sort"
)

func kidsWithCandies(candies []int, extraCandies int) (res []bool) {
	x := make([]int, len(candies))
	copy(x, candies)
	sort.Ints(x)
	maxn := x[len(x)-1]
	for _, i := range candies {
		if i+extraCandies >= maxn {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return
}

func main() {
	fmt.Println(kidsWithCandies([]int{2, 3, 5, 1, 3}, 3))
}
