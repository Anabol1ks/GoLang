package main

import (
	"fmt"
	"strings"
)

func solve(slice []string) (res []int) {
	for _, str := range slice {
		str = strings.ToLower(str)
		count := 0
		for i, ch := range str {
			if i+1 == int(ch)-96 {
				count++
			}
		}
		res = append(res, count)
	}
	return
}

func main() {
	fmt.Println(solve([]string{"abode", "ABc", "xyzD"}))
}
