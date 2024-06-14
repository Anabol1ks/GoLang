package main

import (
	"fmt"
	"strings"
)

func CountDeafRats(town string) int {
	P, c := 0, 0
	left := "O~"
	right := "~O"
	town = strings.ReplaceAll(town, " ", "")
	for ind, i := range town {
		if string(i) == "P" {
			P = ind
		}
	}
	for i := P + 1; i < len(town); i += 2 {
		if string(town[i:i+2]) == right {
			c++
		}
	}
	for i := P - 2; i >= 0; i -= 2 {
		if string(town[i:i+2]) == left {
			c++
		}
	}
	return c
}

func main() {
	fmt.Println(CountDeafRats("~O ~O~O~O~O~O~O~O~OO~O~ O~~O~O~O~OO~~O~O~O~O~O~OO~  ~O~OO~O~~O~O~O~O~O~O ~O~O~OO~~O~O~OO~~O~O~O~O~O~O~O~O~OO~~O~O~O~O~O~O  ~O ~O~O~O~O~O~OPO~~O O~O~O~O~ O~O~O~O~O~~O "))
}
