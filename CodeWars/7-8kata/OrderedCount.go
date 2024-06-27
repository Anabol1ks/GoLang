package main

import "fmt"

// Use the preloaded Tuple struct as return type
type Tuple struct {
	Char  rune
	Count int
}

func OrderedCount(text string) (res []Tuple) {
	if len(text) == 0 {
		return []Tuple{}
	}
	word := make(map[rune]int)
	for _, w := range text {
		word[w]++
	}
	for _, w := range text {
		_, b := word[w]
		if b {
			d := Tuple{
				w,
				word[w],
			}
			res = append(res, d)
			delete(word, w)
		}
	}
	return
}

func main() {
	fmt.Println(OrderedCount(""))
}
