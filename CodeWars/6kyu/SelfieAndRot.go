package main

import (
	"fmt"
	"strings"
)

func Rot(s string) (res string) {
	spl := strings.Split(s, "\n")
	for i := len(spl) - 1; i >= 0; i-- {
		str := spl[i]
		for j := len(str) - 1; j >= 0; j-- {
			res += string(str[j])
		}
		res += "\n"
	}
	return res[:len(res)-1]
}

func SelfieAndRot(s string) (res string) {
	spl := strings.Split(s, "\n")
	for _, str := range spl {
		res += str + strings.Repeat(".", len(str)) + "\n"
	}
	res = res[:len(res)-1]
	res += "\n" + Rot(res)
	return
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

func main() {
	fmt.Println(SelfieAndRot("abcd\nefgh\nijkl\nmnop"))
}
