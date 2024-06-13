package main

import "fmt"

func SumCubes(n int) (res int) {
	for i := 1; i <= n; i++ {
		res += i * i * i
	}
	return
}

func main() {
	fmt.Println(SumCubes(2))
}
