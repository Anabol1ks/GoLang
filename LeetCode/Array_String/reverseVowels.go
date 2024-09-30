package main

import (
	"fmt"
)

func reverseVowels(s string) string {
	pos := []int{}
	v := []rune{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for i, ch := range s {
		for _, b := range v {
			if ch == b {
				pos = append(pos, i)
			}
		}
	}
	r := []rune(s)
	l := len(pos)
	for i := 0; i < l/2; i++ {
		r[pos[i]], r[pos[l-i-1]] = r[pos[l-i-1]], r[pos[i]]
	}

	return string(r)
}

func main() {
	fmt.Println(reverseVowels("leetcode"))
}
