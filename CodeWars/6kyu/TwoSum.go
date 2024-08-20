package main

import "fmt"

func TwoSum(numbers []int, target int) [2]int {
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == target {
				return [2]int{i, j}
			}
		}
	}

	return [2]int{}
}

func main() {
	fmt.Println(TwoSum([]int{1, 2, 3}, 4))
}
