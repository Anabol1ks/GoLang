package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var r, c int
	fmt.Fscan(in, &r, &c)
	arr := make([][]string, r)
	for i := range arr {
		arr[i] = make([]string, c)
	}

	for i := 0; i < r; i++ {
		var line string
		fmt.Fscan(in, &line)
		for j := 0; j < c && j < len(line); j++ {
			arr[i][j] = string(line[j])
		}
	}

	words := []string{}

	for i := 0; i < r; i++ {
		word := ""
		for j := 0; j < c; j++ {
			if arr[i][j] == "#" {
				if len(word) > 1 {
					words = append(words, word)
				}
				word = ""
			} else {
				word += arr[i][j]
			}
		}
		if len(word) > 1 {
			words = append(words, word)
		}
	}

	for j := 0; j < c; j++ {
		word := ""
		for i := 0; i < r; i++ {
			if arr[i][j] == "#" {
				if len(word) > 1 {
					words = append(words, word)
				}
				word = ""
			} else {
				word += arr[i][j]
			}
		}
		if len(word) > 1 {
			words = append(words, word)
		}
	}
	sort.Strings(words)
	fmt.Fprintln(out, words[0])
}
