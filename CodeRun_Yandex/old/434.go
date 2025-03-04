package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var in = bufio.NewReader(os.Stdin)
var out = bufio.NewWriter(os.Stdout)

func razbiv(n, max_v int, cur_sum []int) {
	if n == 0 {
		fmt.Println(strings.Join(intSliceToStringSlice(cur_sum), " + "))
		return
	}
	for i := 1; i <= max_v; i++ {
		if i <= n {
			razbiv(n-i, i, append(cur_sum, i))
		}
	}
}
func intSliceToStringSlice(slice []int) []string {
	strSlice := make([]string, len(slice))
	for i, val := range slice {
		strSlice[i] = fmt.Sprintf("%d", val)
	}
	return strSlice
}

func start(n int) {
	razbiv(n, n, []int{})
}

func main() {
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	start(n)
}
