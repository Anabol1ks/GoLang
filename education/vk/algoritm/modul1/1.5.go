package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	minT := 10000.00
	var a, b int
	for i := 0; i < len(arr)-1; i++ {
		temp := math.Abs(float64(arr[i] - arr[i+1]))
		if temp < minT {
			minT = temp
			a, b = arr[i], arr[i+1]
		}
	}
	fmt.Println(a, b)
}
