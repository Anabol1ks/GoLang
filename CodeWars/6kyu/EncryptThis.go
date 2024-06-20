package main

import (
	"fmt"
	"strconv"
	"strings"
)

func EncryptThis(text string) (res string) {
	sp := strings.Split(text, " ")
	for n, i := range sp {
		for j, _ := range i {
			if j == 0 {
				res += strconv.Itoa(int(rune(i[j])))
				continue
			}
			if j == 1 {
				res += string(i[len(i)-1])
				continue
			}
			if j == len(i)-1 {
				res += string(i[1])
				continue
			}
			res += string(i[j])
		}
		if n != len(sp)-1 {
			res += " "
		}
	}
	return
}

func main() {
	fmt.Println(EncryptThis("Hello Word"))
}
