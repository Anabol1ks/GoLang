package main

import (
	"fmt"
)

func SeriesSum(n int) string {
	if n == 0 {
		return "0.00"
	}
	var x, d float64 = 1, 0
	for i := 1; i < n; i++ {
		if n > 1 {
			d = 1.00 / (4 + 3*float64(i-1))
			x = x + d
		}

	}
	return fmt.Sprintf("%.2f", x)
}

func main() {
	fmt.Println(SeriesSum(5))
}
