package main

import (
	"fmt"
	"strings"
)

func StringToArray(str string) []string {
	s := strings.Split(str, " ")
	return s
}

func main() {
	s := StringToArray("Robin Singh")
	fmt.Printf("%q", s)
}
