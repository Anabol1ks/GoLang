package main

import (
	"fmt"
	"regexp"
)

func main() {
	var str, res string
	fmt.Scan(&str)
	count := 0
	re := regexp.MustCompile(`[a-zA-Zа-яА-Я]+`)
	words := re.FindAllString(str, -1)
	maxLen := 0

	for _, word := range words {
		if len(word) > maxLen {
			maxLen = len(word)
		}
	}
	len := maxLen * 3
	for _, i := range str {
		if count == len {
			count = 0
			res += "\n"
		}
		res += string(i)
		count++
		if string(i) == "," {
			res += " "
			count++
		}
	}
	fmt.Println(len, maxLen)
	fmt.Println("=============")
	fmt.Println(str, res)
}
