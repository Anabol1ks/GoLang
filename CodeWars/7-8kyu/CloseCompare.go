package main

import (
	"fmt"
	"math"
)

func CloseCompare(a, b, margin float64) (res int) {
	if math.Abs(a-b) <= margin {
		res = 0
	} else if a > b {
		res = 1
	} else if a < b {
		res = -1
	}
	return res
}

func main() {
	fmt.Println(CloseCompare(315, 10, 3))
}
