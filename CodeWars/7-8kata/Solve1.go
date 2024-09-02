package main

import "fmt"

func i(a, b int) []int {
	if a == 0 || b == 0 {
		return []int{a, b}
	}
	return ii(a, b)
}

func ii(a, b int) []int {
	if a >= 2*b {
		a = a - 2*b
		return i(a, b)
	}
	return iii(a, b)
}

func iii(a, b int) []int {
	if b >= 2*a {
		b = b - 2*a
		return i(a, b)
	}
	return []int{a, b}
}

func Solve(a, b int) []int {
	return i(a, b)
}

func main() {
	fmt.Println(Solve(6, 19))
}
