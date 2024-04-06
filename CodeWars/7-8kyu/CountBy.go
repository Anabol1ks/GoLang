package main

import "fmt"

func CountBy(x, n int) (m []int) {
	for i := x; i <= x*n; i += x {
		m = append(m, i)
	}
	return m
}

func main() {
	fmt.Println(CountBy(3, 5))
}
