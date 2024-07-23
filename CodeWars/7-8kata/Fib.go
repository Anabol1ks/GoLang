package main

import "fmt"

func Fib(n int) int {
	arr := []int{0, 1}
	for i := len(arr) - 1; i <= n; i++ {
		arr = append(arr, arr[i]+arr[i-1])
	}
	return arr[n]
}

func main() {
	fmt.Println(Fib(5))
}
