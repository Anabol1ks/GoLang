package main

import (
	"fmt"
	"strings"
)

func FindShort(s string) (minLen int) {
	minLen = 100000
	spl := strings.Split(s, " ")
	for _, str := range spl {
		if len(str) < minLen {
			minLen = len(str)
		}
	}
	return
}

func main() {
	fmt.Println(FindShort("turns out random test cases are easier than writing out basic ones"))
}
