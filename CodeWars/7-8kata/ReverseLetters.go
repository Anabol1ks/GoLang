package main

import (
	"fmt"
	"unicode"
)

func ReverseLetters(s string) (res string) {
	var result []rune
	for _, r := range s {
		if unicode.IsLetter(r) {
			result = append(result, r)
		}
	}
	str := string(result)
	for _, i := range str {
		res = string(i) + res
	}
	return
}

func main() {
	fmt.Println(ReverseLetters("kap!iba2ra"))
}
