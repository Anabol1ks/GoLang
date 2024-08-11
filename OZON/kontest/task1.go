package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var num int
		fmt.Fscan(in, &num)
		if num < 10 {
			fmt.Fprintln(out, 0)
			continue
		}
		strNum := strconv.Itoa(num)
		maxNum := 0
		for j := 0; j < len(strNum); j++ {
			newNumStr := strNum[:j] + strNum[j+1:]
			newNum, _ := strconv.Atoi(newNumStr)
			if newNum > maxNum {
				maxNum = newNum
			}
		}

		fmt.Fprintln(out, maxNum)
	}
}
