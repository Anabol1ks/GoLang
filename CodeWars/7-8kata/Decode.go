package main

import (
	"fmt"
	"strconv"
)

func Decode(code []int, key int) (res string) {
	bl := make(map[int]string)
	n := 1
	for i := 'a'; i <= 'z'; i++ {
		bl[n] = string(i)
		n++
	}
	num := []int{}
	s := strconv.Itoa(key)
	for _, r := range s {
		d, _ := strconv.Atoi(string(r))
		num = append(num, d)
	}
	for i, n := range code {
		w := (i + 1) % len(s)
		if w == 0 {
			w = len(s)
		}
		it := n - num[w-1]
		res += bl[it]
	}
	return
}

func main() {
	fmt.Println(Decode([]int{20, 12, 18, 30, 21}, 1939))
}
