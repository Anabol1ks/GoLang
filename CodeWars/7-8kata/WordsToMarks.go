package main

import "fmt"

func WordsToMarks(s string) int {
	sigma := make(map[string]int)
	c := 1
	for i := 'a'; i <= 'z'; i++ {
		sigma[string(i)] = c
		c++
	}
	sum := 0
	for _, v := range s {
		sum += sigma[string(v)]
	}
	return sum
}

func main() {
	fmt.Println(WordsToMarks("love"))
}
