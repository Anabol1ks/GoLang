package main

import "fmt"

func RoundToNext5(n int) int {
	if n%5 == 0 {
		return n
	}
	if n < 0 {
		return n - n%5
	}
	return n + (5 - n%5)
}

func main() {
	fmt.Println(RoundToNext5(-2))
}
