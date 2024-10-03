package main

import "fmt"

func isSubsequence(s string, t string) bool {
	c := 0
	id := -1
	for _, s1 := range s {
		for i, t1 := range t {
			if s1 == t1 && id < i {
				id = i
				c++
				break
			}
		}
	}
	if c == len(s) {
		return true
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("ab", "baab"))
}
