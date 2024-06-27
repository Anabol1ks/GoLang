package main

import (
	"fmt"
	"strconv"
)

func FreqSeq(str string, sep string) (res string) {
	counts := make(map[rune]int)
	for _, s := range str {
		counts[s]++
	}
	for i, s := range str {
		if i != len(str)-1 {
			res += strconv.Itoa(counts[s]) + sep
		} else {
			res += strconv.Itoa(counts[s])
		}
	}
	return res
}

func main() {
	fmt.Println(FreqSeq("hello world", "-"))
}
