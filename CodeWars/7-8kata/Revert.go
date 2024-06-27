package main

import "fmt"

func ReverseSeq(n int) []int {
	d := []int{}
	for i := n; i > 0; i-- {
		d = append(d, i)
	}
	return d
}

func main() {
	n := 101
	fmt.Println(ReverseSeq(n))
}
