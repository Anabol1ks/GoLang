package main

import "fmt"

func main() {
	var n, target int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	left, right := 0, len(arr)-1
	fmt.Scan(&target)
	fmt.Println(binarySearchPlus(arr, left, right, target))
}

func binarySearchPlus(arr []int, left, right, target int) int {
	if left > right {
		return -1
	}
	middle := (left + right) / 2
	if arr[middle] == target {
		return middle
	}
	if arr[middle] > target {
		return binarySearchPlus(arr, left, middle-1, target)
	} else {
		if middle+1 == right {
			return right
		}
		return binarySearchPlus(arr, middle+1, right, target)
	}
}
