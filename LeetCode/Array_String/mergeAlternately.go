package main

import "fmt"

func mergeAlternately(word1 string, word2 string) (res string) {
	if len(word1) > len(word2) {
		for i, _ := range word2 {
			res += string(word1[i]) + string(word2[i])
		}
		res += string(word1[len(word2):])
	} else {
		for i, _ := range word1 {
			res += string(word1[i]) + string(word2[i])
		}
		res += string(word2[len(word1):])
	}
	return
}

func main() {
	word1 := "ab"
	word2 := "pqrs"

	fmt.Println(mergeAlternately(word1, word2))
}
