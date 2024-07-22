package main

import (
	"fmt"
)

type WordLen struct {
	Word string
	Lens int
}

func LongestConsec(strarr []string, k int) string {
	if k > len(strarr) {
		return ""
	}
	wordlen := make([]WordLen, len(strarr))
	for i, w := range strarr {
		wordlen[i] = WordLen{Word: w, Lens: len(w)}
	}
	sum, max := 0, 0
	str := ""
	for i := 0; i <= len(wordlen)-k; i++ {
		d := wordlen[i : i+k]
		for j := 0; j < len(d); j++ {
			sum += d[j].Lens
		}
		if sum > max {
			str = ""
			max = sum
			for j := 0; j < len(d); j++ {
				str += d[j].Word
			}
		}
		sum = 0
	}
	return str
}

func main() {
	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
}
