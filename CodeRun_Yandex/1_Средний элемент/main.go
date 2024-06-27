package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, b, c int
	fmt.Scanln(&a, &b, &c)
	r := make([]int, 0)
	r = append(r, a, b, c)
	sort.Ints(r)
	fmt.Println(r[1])
}
