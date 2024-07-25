package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t, n int
	var p float64
	fmt.Fscanln(in, &t)
	for i := 0; i < t; i++ {
		fmt.Fscan(in, &n, &p)
		var g, raz float64
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &g)
			k := g * p / 100
			raz += k - float64(int(k))
		}
		raz = math.Round(raz*100) / 100
		fmt.Fprintf(out, "%.2f\n", raz)
	}
}
