package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	al, _ := in.ReadString('\n')
	bukv, _ := in.ReadString('\n')
	var maxLen int
	fmt.Fscanf(in, "%d", &maxLen)

	al = strings.TrimSpace(al)
	bukv = strings.TrimSpace(bukv)

	bukvChars := make(map[rune]bool)
	for _, ch := range bukv {
		bukvChars[ch] = true
	}

	wStart := 0
	wCount := make(map[rune]int)
	bukvCount := len(bukvChars)
	foundCount := 0
	result := ""
	n := len(al)

	for wEnd := 0; wEnd < n; wEnd++ {
		charEnd := rune(al[wEnd])
		if _, exists := bukvChars[charEnd]; exists {
			wCount[charEnd]++
			if wCount[charEnd] == 1 {
				foundCount++
			}
		}

		for wEnd-wStart+1 > maxLen {
			charStart := rune(al[wStart])
			if _, exists := bukvChars[charStart]; exists {
				wCount[charStart]--
				if wCount[charStart] == 0 {
					foundCount--
				}
			}
			wStart++
		}

		if wEnd-wStart+1 <= maxLen && foundCount == bukvCount {
			curW := al[wStart : wEnd+1]
			if result == "" || (strings.Index(al, curW) > strings.Index(al, result)) || len(curW) > len(result) {
				result = curW
			}
		}
	}

	if result == "" {
		fmt.Fprintln(out, -1)
	} else {
		fmt.Fprintln(out, result)
	}
}
