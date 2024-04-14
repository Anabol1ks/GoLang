package main

import "fmt"

func Divisors(n int) (count int) {
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return
}

func main() {
	fmt.Println(Divisors(54))
}
