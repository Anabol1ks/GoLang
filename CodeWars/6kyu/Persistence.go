package main

import (
	"fmt"
	"strconv"
)

func Persistence(n int) (c int) {
	for n >= 10 {
		r := 1
		str := strconv.Itoa(n)
		dig := make([]int, len(str))
		for i, char := range str {
			dig[i] = int(char - '0')
		}
		for _, num := range dig {
			r *= num
			n = r
		}
		c++
	}
	return
}

func main() {
	fmt.Println(Persistence(39))
}
