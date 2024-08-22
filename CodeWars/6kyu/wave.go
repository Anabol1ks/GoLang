package main

import (
	"fmt"
	"strings"
)

func wave(words string) (res []string) {
	for i, ch := range words {
		if string(ch) == " " {
			continue
		} else {
			str := words[:i] + strings.ToUpper(string(ch)) + words[i+1:]
			res = append(res, str)
		}
	}
	return
}

func main() {
	fmt.Println(wave(" a aa  k"))
}
