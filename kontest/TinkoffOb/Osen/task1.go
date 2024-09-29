package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var str, res string

	fmt.Fscanln(in, &str)
	sl := strings.Split(str, ",")
	for _, s := range sl {
		sl1 := strings.Split(s, "-")
		if len(sl1) > 1 {
			n1, _ := strconv.Atoi(sl1[0])
			n2, _ := strconv.Atoi(sl1[1])
			for i := n1; i <= n2; i++ {
				res += strconv.Itoa(i) + " "
			}
		} else {
			res += sl1[0]
		}
	}
	fmt.Fprintln(out, res)
}
