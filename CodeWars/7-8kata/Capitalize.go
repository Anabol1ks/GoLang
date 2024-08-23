package main

import (
	"fmt"
	"strings"
)

func Capitalize(st string) (res []string) {
	var str1, str2 string
	for i, ch := range st {
		buk := string(ch)
		if (i+1)%2 == 0 {
			str1 += strings.ToUpper(buk)
			str2 += buk
		} else {
			str2 += strings.ToUpper(buk)
			str1 += buk
		}
	}
	res = append(res, str2, str1)
	return
}

func main() {
	fmt.Println(Capitalize("abcdefg"))
}
