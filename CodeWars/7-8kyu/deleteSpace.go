package main

import (
	"fmt"
	"strings"
)

func NoSpace(word string) string {
	return strings.ReplaceAll(word, " ", "")
}

func main() {
	str := "8 j 8   mBliB8g  imjB8B8  jl  B"
	fmt.Println(NoSpace(str))
}
