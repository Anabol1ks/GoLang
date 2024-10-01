package main

import "fmt"

func moveZeroes(nums []int) []int {
	nonZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums, "----")
		if nums[i] != 0 {
			nums[nonZeroIndex] = nums[i]
			nonZeroIndex++
		}
	}

	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
	return nums
}

func main() {
	fmt.Println(moveZeroes([]int{0, 1, 0, 3, 12, 0}))
}
