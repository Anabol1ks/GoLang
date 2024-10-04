package main

import (
	"fmt"
)

func increasingTriplet(nums []int) bool {
	first, second := int(^uint(0)>>1), int(^uint(0)>>1)
	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(increasingTriplet([]int{20, 100, 10, 12, 5, 13}))
}
