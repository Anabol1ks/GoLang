package main

import "fmt"

func Maps(x []int) (d []int) {
	for i, _ := range x {
		d = append(d, x[i]*2)
	}
	return
}

func main() {
	fmt.Println(Maps([]int{1, 2, 3}))
}
