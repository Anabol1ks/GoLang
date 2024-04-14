package main

import (
	"fmt"
	"strings"
)

func solve(str string) string {
	lower, up := 0, 0
	for _, i := range str {
		for j := 'A'; j <= 'Z'; j++ {
			if string(i) == string(j) {
				up++
			}
		}
		for f := 'a'; f <= 'z'; f++ {
			if string(i) == string(f) {
				lower++
			}
		}
	}
	if lower >= up {
		return strings.ToLower(str)
	}
	return strings.ToUpper(str)
}

func main() {
	fmt.Println(solve("cODE"))
}
