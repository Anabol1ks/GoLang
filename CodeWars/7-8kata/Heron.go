package main

import (
	"fmt"
	"math"
)

func Heron(a, b, c float64) (area float64) {
	s := (a + b + c) / 2
	area = math.Sqrt(s * (s - a) * (s - b) * (s - c))
	return
}

func mains() {
	fmt.Println(Heron(3.0, 4.0, 5.0))
}
