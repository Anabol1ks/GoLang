package main

import (
	"fmt"
	"strings"
)

func SortVowels(s string) (res string) {
	gl := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
	}
	s_l := strings.ToLower(s)
	for i, _ := range s {
		if gl[string(s_l[i])] == true {
			res += "|" + string(s[i]) + "\n"
		} else {
			res += string(s[i]) + "|" + "\n"
		}
	}
	if len(res) > 0 {
		res = res[:len(res)-1]
	}
	return
}

func main() {
	fmt.Println(SortVowels("CodewArs"))
}
