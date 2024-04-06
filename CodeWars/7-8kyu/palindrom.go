package main

import (
	"fmt"
	"strings"
)

func IsPalindrome(str string) bool {
	str = strings.ToLower(str)
	lend := len(str)
	for i := 0; i < lend/2; i++ {
		if str[i] != str[lend-1-i] {
			return false
		}
	}
	return true
}

func main() {
	d := "addA"
	fmt.Println(IsPalindrome(d))
}
