package main

import "fmt"

func MaximumSubarraySum(numbers []int) int {
	lArr := len(numbers)
	cNet := 0
	maxSum := 0
	for i := 0; i < lArr; i++ {
		sum := 0
		for j := i; j < lArr; j++ {
			if numbers[i] < 0 {
				cNet++
			}
			sum += numbers[j]
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	if lArr == 0 || cNet == lArr {
		return 0
	}
	return maxSum
}

func main() {
	fmt.Println(MaximumSubarraySum([]int{-2, -1, -4, -4, 8}))
}
