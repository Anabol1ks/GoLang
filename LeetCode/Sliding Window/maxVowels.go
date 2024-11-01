package main

import (
	"fmt"
)

func maxVowels(s string, k int) int {
	isVowel := func(ch byte) bool {
		return ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u'
	}

	c, maxc := 0, 0
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			c++
		}
	}
	maxc = c
	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			c++
		}
		if isVowel(s[i-k]) {
			c--
		}
		if c > maxc {
			maxc = c
		}
	}

	return maxc
}

func main() {
	fmt.Println(maxVowels("weallloveyou", 7))
}
