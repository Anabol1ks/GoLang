package main

import (
	"fmt"
)

func SortMyString(s string) string {
	d, w := "", ""
	for i, v := range s {
		if i%2 == 0 {
			d += string(v)
		}
		if i%2 == 1 {
			w += string(v)
		}
	}
	s = d + " " + w
	return s
}

func main() {
	fmt.Printf(SortMyString("CodeWars"))
}
