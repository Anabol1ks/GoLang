package main

import (
	"fmt"
	"math"
)

func WallPaper(l, w, h float64) string {
	if l == 0 || w == 0 || h == 0 {
		return "zero"
	}
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen", "twenty"}
	pl := (l*h*2 + (w * h * 2)) * 1.15 / 5.2
	res := math.Ceil(pl)

	return numbers[int(res)]
}

func main() {
	fmt.Println(WallPaper(6.3, 5.8, 3.13))
}
