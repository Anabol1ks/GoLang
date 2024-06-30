package main

import (
	"fmt"
	"strconv"
)

func Encode(str string, key int) (res []int) {
	bl := make(map[string]int)
	n := 1
	for i := 'a'; i <= 'z'; i++ {
		bl[string(i)] = n
		n++
	}
	num := []int{}
	s := strconv.Itoa(key)
	for _, r := range s {
		d, _ := strconv.Atoi(string(r))
		num = append(num, d)
	}
	for i, ch := range str {
		w := (i + 1) % len(s)
		if w == 0 {
			w = len(s)
		}
		it := bl[string(ch)] + num[w-1]
		res = append(res, it)
	}
	return
}

func main() {
	fmt.Println(Encode("masterpiece", 1939))
}
