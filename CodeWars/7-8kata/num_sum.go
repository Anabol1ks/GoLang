package main

import "fmt"

func Summation(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return
}

func main() {
	fmt.Println(Summation(3))
}
