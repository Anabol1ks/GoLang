package main

import (
	"fmt"
	"strings"
)

func solution(str, ending string) (H bool) {
	if strings.HasSuffix(str, ending) {
		H = true
	}
	return
}

func main() {

	fmt.Println(solution("abc", "bc"))
}
