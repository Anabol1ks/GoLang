package main

import "fmt"

func MaxBall(v0 int) int {
	v := float64(v0) / 3.6
	g := 9.81
	count := 0
	var h1, h_max float64
	for t := 0.1; ; t += 0.1 {
		h := v*t - 0.5*g*t*t
		h1 = h
		if h1 < h_max {
			break
		} else {
			h_max = h1
		}
		count++
	}
	return count
}

func main() {
	fmt.Println(MaxBall(99))
}
