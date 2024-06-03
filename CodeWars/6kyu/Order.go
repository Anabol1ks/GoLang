package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Order(sentence string) string {
	new_sentence := strings.Split(sentence, " ")
	sentence = ""
	for n := 1; n <= len(new_sentence); n++ {
		for _, i := range new_sentence {
			for _, j := range string(i) {
				s, _ := strconv.Atoi(string(j))
				for s == n {
					sentence += i
					if n != len(new_sentence) {
						sentence += " "
					}
					break
				}
			}
		}
	}
	return sentence
}

func main() {
	fmt.Println(Order("is2 Thi1s T4est 3a"))
}
