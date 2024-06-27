package main

import (
	"fmt"
)

func Xor(a, b bool) bool {
	if a == false && b == true || a == true && b == false {
		return true
	}
	return false
}

func main() {
	fmt.Println(Xor(true, true))
}
