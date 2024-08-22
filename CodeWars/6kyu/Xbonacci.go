package main

import "fmt"

func Xbonacci(signature []int, n int) (res []int) {
	larr := len(signature)
	res = append(res, signature...)
	for i := larr; i < n; i++ {
		sum := 0
		for j := i - larr; j < i; j++ {
			sum += res[j]
		}
		res = append(res, sum)
	}

	if n < larr {
		return res[:n]
	}
	return
}

func main() {
	fmt.Println(Xbonacci([]int{1, 1, 1, 1}, 10))
}
