package main

import "fmt"

func CheckForFactor(base int, factor int) bool {
	if base%factor == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(CheckForFactor(12, 9))
}
