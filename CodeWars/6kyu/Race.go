package main

import (
	"fmt"
	"math"
)

func Race(v1, v2, g int) [3]int {
	if v1 >= v2 {
		return [3]int{-1, -1, -1}
	}
	t := math.Abs(float64(g) / float64(v1-v2))
	h := int(t)
	m := int((t - float64(h)) * 60)
	s := int(((t-float64(h))*60 - float64(m)) * 60)
	return [3]int{h, m, s}
}

func main() {
	fmt.Println(Race(820, 850, 550))
}
