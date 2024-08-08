package main

import (
	"fmt"
	"sort"
)

func TwoToOne(s1 string, s2 string) string {
	ca := make(map[string]bool)
	str := []byte(s1 + s2)
	for _, i := range str {
		ca[string(i)] = true
	}
	res := ""
	for i, _ := range ca {
		if ca[i] == true {
			res += i
		}
	}
	result := []byte(res)
	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	return string(result)
}

func main() {
	fmt.Println(TwoToOne("abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"))
}
