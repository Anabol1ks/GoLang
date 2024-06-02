package main

import "fmt"

func NbYear(p0 int, percent float64, aug int, p int) (c int) {
	for p0 < p {
		p0 = p0 + int(float64(p0)*percent/100) + aug
		c++
	}
	return c
}

func main() {
	fmt.Println(NbYear(1500, 5, 100, 5000))
}
