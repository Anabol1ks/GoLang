package main

import (
	"fmt"
	"math"
)

func FindNextSquare(sq int64) int64 {
	a := float64(sq)
	x := math.Sqrt(a)
	var d int64
	if math.Mod(x, 1) == 0 {
		x++
		for x < x*2 {
			if math.Mod(x, 1) == 0 {
				d = int64(x * x)
				break
			}
		}
	} else {
		d = -1
	}
	return d
}

func main() {
	fmt.Println(FindNextSquare(36))
}
