package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	var maxC int
	for _, i := range arr {
		if i > maxC {
			maxC = i
		}
	}
	fmt.Println(maxC)
}
