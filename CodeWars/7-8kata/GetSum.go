package main

import "fmt"

func GetSum(a, b int) (res int) {
	if a > b {
		a, b = b, a
	}
	for i := a; i <= b; i++ {
		res += i
	}
	return
}

func main() {
	fmt.Println(GetSum(1, 0))
}
