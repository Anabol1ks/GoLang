package main

import "fmt"

func Seven(n int64) (res []int) {
	c := 0
	for n/100 != 0 {
		n2 := n % 10
		n1 := n / 10
		n = n1 - n2*2
		c++
	}
	res = append(res, int(n), c)
	return
}

func main() {
	fmt.Println(Seven(371))
}
