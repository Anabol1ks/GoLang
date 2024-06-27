package main

import (
	"fmt"
	"math"
)

func PowersOfTwo(n int) (res []uint64) {
	for i := 0; i <= n; i++ {
		s := math.Pow(2, float64(i))
		res = append(res, uint64(s))
	}
	return
}

func main() {
	fmt.Println(PowersOfTwo(5))
}
