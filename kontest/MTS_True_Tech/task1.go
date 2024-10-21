package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var a, b, c int
	fmt.Fscan(in, &a)
	fmt.Fscan(in, &b)
	fmt.Fscan(in, &c)
	if c > 100 {
		c = a + (c-100)*b
		fmt.Fprintln(out, c)
	} else {
		fmt.Fprintln(out, a)
	}
}
