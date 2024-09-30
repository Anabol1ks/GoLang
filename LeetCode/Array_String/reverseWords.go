package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	res := ""
	w := strings.Fields(s)
	s = strings.Join(w, " ")
	str := strings.Split(s, " ")
	l := len(str)
	for i := 0; i < l/2; i++ {
		str[i], str[l-i-1] = str[l-i-1], str[i]
	}
	res = strings.Join(str, " ")
	return res
}

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}
