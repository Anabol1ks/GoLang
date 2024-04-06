package main

import (
	"fmt"
	"strings"
)

func Solution(word string) string {
	a := strings.Split(word, "")
	b := []string{}
	for i := 0; i < len(a); i++ {
		b = append(b, a[len(a)-i-1])
	}
	return strings.Join(b, "")
}

func main() {
	str := "word"

	fmt.Println(Solution(str))
}
