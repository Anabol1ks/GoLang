package main

import "fmt"

func Invert(arr []int) []int {
	var new_arr []int
	for i := 0; i < len(arr); i++ {
		new_arr = append(new_arr, arr[i]*-1)
	}
	return new_arr
}

func main() {
	s := []int{1, -2, 3, -4, 5}
	fmt.Println(Invert(s))
}
