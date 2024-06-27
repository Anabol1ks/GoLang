package main

import (
	"fmt"
)

func Gimme(array [3]int) (res int) {
	max, min := -10000, 10000
	for _, i := range array {
		if i > max {
			max = i
		}
		if i < min {
			min = i
		}
	}
	for i, _ := range array {
		if array[i] != max && array[i] != min {
			res = i
		}
	}
	return res
}

func main() {
	fmt.Println(Gimme([3]int{3, 2, 1}))

}
