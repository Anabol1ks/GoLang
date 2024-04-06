package main

import (
	"fmt"
	"strings"
)

func Disemvowel(comment string) string {
	sogl := []string{"a", "e", "u", "i", "o", "A", "E", "U", "I", "O"}
	for _, i := range sogl {
		comment = strings.ReplaceAll(comment, i, "")
	}
	return comment
}

func main() {
	a := "This website is for losers LOL!"
	fmt.Println(Disemvowel(a))
}
