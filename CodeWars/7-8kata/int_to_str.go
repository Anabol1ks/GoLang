package main

import (
	"fmt"
	"strconv"
)

func NumberToString(n int) string {
	return strconv.Itoa(n)
}

func main() {
	n := -123
	fmt.Printf("%q", NumberToString(n))
}
