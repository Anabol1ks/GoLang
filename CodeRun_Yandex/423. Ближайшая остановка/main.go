package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	o := make([]int, 0)
	a := make([]int, 0)
	for i := 0; i < n; i++ {
		var num int
		fmt.Scan(&num)
		o = append(o, num)
	}
	for i := 0; i < k; i++ {
		var num int
		fmt.Scan(&num)
		a = append(a, num)
	}
	fmt.Println()
	bl, idx := math.MaxInt64, math.MaxInt64
	for _, i_a := range a {
		for id_o, i_o := range o {
			r := int(math.Abs(float64(i_a) - float64(i_o)))
			if r < bl {
				bl = r
				idx = id_o
			} else if r == bl {
				if id_o < idx {
					idx = id_o
				}
			}
		}
		fmt.Println(idx + 1)
	}
}
