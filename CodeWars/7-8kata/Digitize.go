package main

import "fmt"

func Digitize(n int) []int {
	r := make([]int, 0)
	if n == 0 {
		return []int{0}
	}
	for n > 0 {
		m := n % 10
		r = append(r, m)
		n = n / 10
	}
	return r
}

func main() {
	fmt.Println(Digitize(35231))
}
