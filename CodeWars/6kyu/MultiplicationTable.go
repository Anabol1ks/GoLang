package main

import "fmt"

func MultiplicationTable(size int) [][]int {
	d := make([][]int, size)
	for i := range d {
		d[i] = make([]int, size)
	}
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			d[i-1][j-1] = i * j
		}
	}
	return d
}

func main() {
	fmt.Println(MultiplicationTable(3))
}
