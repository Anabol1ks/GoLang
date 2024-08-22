package main

import (
	"fmt"
	"strings"
)

func bandNameGenerator(word string) string {
	if word[0] == word[len(word)-1] {
		return strings.Title(word) + word[1:]
	}
	return "The " + strings.Title(word)
}

func main() {
	fmt.Println(bandNameGenerator("alaskd"))
}
