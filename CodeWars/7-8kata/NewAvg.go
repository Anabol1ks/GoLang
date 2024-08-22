package main

import (
	"fmt"
	"math"
)

func NewAvg(arr []float64, navg float64) int64 {
	var sum float64
	l := len(arr) + 1
	for _, i := range arr {
		sum += i
	}
	res := float64(l)*navg - sum
	if res > 0 {
		return int64(math.Ceil(res))
	}
	return -1
}

func main() {
	fmt.Println(NewAvg([]float64{14.0, 30.0, 5.0, 7.0, 9.0, 11.0, 15.0}, 30))
}
