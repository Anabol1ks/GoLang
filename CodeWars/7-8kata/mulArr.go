package main

import "fmt"

func Grow(arr []int) int {
	mul := 1
	for i := 0; i < len(arr); i++ {
		mul *= arr[i]
	}
	return mul
}

func main() {
	m := Grow([]int{1, 2, 3, 4})
	fmt.Println(m)
}
