package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	data := make([]int, n)
	for i := range data {
		fmt.Scan(&data[i])
	}

	var target int
	fmt.Scan(&target)

	fmt.Println(exponentialSearch(data, target))
}

func exponentialSearch(data []int, target int) (int, int) {
	border := 1
	for border < len(data)-1 {
		if data[border] < target {
			border *= 2
		} else if data[border] >= target {
			return border / 2, border
		}
	}
	return 0, 0
}
