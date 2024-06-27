package main

import (
	"fmt"
	"strings"
)

func ToJadenCase(str string) (res string) {
	str_n := strings.Split(str, " ")
	for n, i := range str_n {
		res += strings.ToUpper(string(i[0]))
		for s := 1; s < len(i); s++ {
			res += string(i[s])
		}
		if n != len(str_n)-1 {
			res += " "
		}
	}
	return res
}

func main() {
	fmt.Println(ToJadenCase("most trees are blue"))
}
