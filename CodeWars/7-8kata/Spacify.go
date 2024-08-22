package main

import (
	"fmt"
	"strings"
)

func Spacify(s string) string {
	if len(s) > 0 {
		res := strings.ReplaceAll(s, "", " ")
		return res[1 : len(res)-1]
	}
	return ""
}

func main() {
	fmt.Println(Spacify("hello world"))
	fmt.Println(Spacify(""))
}
