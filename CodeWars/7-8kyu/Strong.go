package main

import (
	"fmt"
	"strconv"
)

func Strong(n int) string {
	s := strconv.Itoa(n)
	sum := 0
	for _, i := range s {
		newN, _ := strconv.Atoi(string(i))
		c := 1
		for j := 1; j <= newN; j++ {
			c *= j
		}
		sum += c
	}
	if sum == n {
		return "STRONG!!!!"
	}
	return "Not Strong !!"
}

func main() {
	fmt.Println(Strong(145))
}
