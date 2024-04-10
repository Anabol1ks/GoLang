package main

import (
	"fmt"
	"strings"
)

func duplicate_count(s1 string) int {
	w := strings.ToLower(s1)
	count := 0
	checked := make(map[rune]bool)
	for i, char := range w {
		if checked[char] {
			continue
		}
		for j := i + 1; j < len(w); j++ {
			if char == rune(w[j]) {
				count++
				checked[char] = true
				break
			}
		}
	}
	return count
}

func main() {
	fmt.Println(duplicate_count("indivisibility"))
}
