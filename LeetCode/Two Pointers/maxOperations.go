package main

import "fmt"

func maxOperations(nums []int, k int) (res int) {
	counts := make(map[int]int)
	for _, num := range nums {
		complement := k - num
		if counts[complement] > 0 {
			res++
			counts[complement]--
		} else {
			counts[num]++
		}
	}
	return
}

func main() {
	fmt.Println(maxOperations([]int{1, 2, 3, 4}, 5))
}
