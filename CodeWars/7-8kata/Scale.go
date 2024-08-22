package main

import (
	"fmt"
	"strings"
)

func scaleK(s string, k int) (res string) {
	str := strings.Split(s, "\n")
	for _, strock := range str {
		for _, i := range strock {
			res += strings.Repeat(string(i), k)
		}
		res += "\n"
	}
	res = res[:len(res)-1]
	return
}

func scaleV(s string, n int) (res string) {
	str := strings.Split(s, "\n")
	for _, strock := range str {
		res += strings.Repeat(strock+"\n", n)
	}
	res = res[:len(res)-1]
	return
}

func Scale(s string, k, n int) (res string) {
	if len(s) == 0 {
		return ""
	}
	res = scaleV(scaleK(s, k), n)
	return
}

func main() {
	fmt.Println(Scale("abcd\nefgh\nijkl\nmnop", 2, 3))
}
