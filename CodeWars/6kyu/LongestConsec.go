package main

import (
	"fmt"
	"sort"
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
	sort.Slice(wordlen, func(i, j int) bool {
		return wordlen[i].Lens > wordlen[j].Lens
	})
	longestWords := make([]string, 0, k)

	for i := 0; i < k && i < len(wordlen); i++ {
		longestWords = append(longestWords, wordlen[i].Word)
	}
	str := ""
	for _, s := range longestWords {
		str += s
	}
	return str
}

func main() {
	fmt.Println(LongestConsec([]string{"zone", "abigail", "theta", "form", "libe", "zas"}, 2))
}
