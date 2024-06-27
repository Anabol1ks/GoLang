package main

import (
	"fmt"
	"strings"
)

func VertMirror(s string) (new_string string) {
	s_plus := strings.Split(s, "\n")
	for a, i := range s_plus {
		for w := len(i) - 1; w >= 0; w-- {
			new_string += string(i[w])
		}
		if a != len(s_plus)-1 {
			new_string += "\n"
		}
	}
	return
}

func HorMirror(s string) (new_string string) {
	s_plus := strings.Split(s, "\n")
	for i := len(s_plus) - 1; i >= 0; i-- {
		new_string += s_plus[i]
		if i != 0 {
			new_string += "\n"
		}
	}
	return
}

type FParam func(string) string

func Oper(f FParam, x string) string {
	return f(x)
}

func main() {
	s := "abcd\nefgh\nijkl\nmnop"
	fmt.Println(Oper(HorMirror, s))
}
