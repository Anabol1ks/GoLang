package main

import (
	"fmt"
	"strconv"
)

func Derive(coefficient, exponent int) string {
	return strconv.Itoa(coefficient*exponent) + "x" + "^" + strconv.Itoa(exponent-1)
}

func main() {
	fmt.Println(Derive(7, 8))
}
