package main

import "fmt"

func findMaxAverage(nums []int, k int) (ma float64) {
	l := len(nums)
	if l > 1 {
		for i := 0; i < l-3; i++ {
			a := 0.0
			sum := 0
			for j := i; j < i+k; j++ {
				sum += nums[j]
			}
			a = float64(sum) / float64(k)
			if a > ma {
				ma = a
			}
			a = 0
		}
	} else {
		ma = float64(nums[0])
	}
	return ma
}

func main() {
	fmt.Println(findMaxAverage([]int{0, 1, 1, 3, 3}, 4))
}
