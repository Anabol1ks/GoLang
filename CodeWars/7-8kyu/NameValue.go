package main

import (
	"fmt"
)

func NameValue(my_list []string) (res []int) {
	al := make(map[string]int)
	n := 1
	for i := 'a'; i <= 'z'; i++ {
		al[string(i)] = n
		n++
	}
	for i, s := range my_list {
		sum := 0
		for _, o := range s {
			sum += al[string(o)]
		}
		res = append(res, sum*(i+1))
	}
	return
}

func main() {
	fmt.Println(NameValue([]string{"abc", "abc", "abc", "abc"}))
}
