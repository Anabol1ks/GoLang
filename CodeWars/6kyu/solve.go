package main

import (
	"fmt"
	"strings"
)

func solve(str string) int {
	alpho := map[string]int{}
	alpho_int := 1
	gl := "aeiou"
	for i := 'a'; i <= 'z'; i++ {
		alpho[string(i)] = alpho_int
		alpho_int++
	}
	var count, maxC int
	for _, ch := range str {
		if strings.ContainsRune(gl, ch) {
			count = 0
		} else {
			count += alpho[string(ch)]
			if count > maxC {
				maxC = count
			}
		}
	}

	return maxC
}

func main() {
	fmt.Println(solve("abababababfapeifapefijaefaepfjavnefjnfbhwyfnjsifjapnes"))
}
