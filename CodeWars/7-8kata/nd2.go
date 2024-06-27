package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	n := 25
	d := 1
	count := 0
	for i := 0; i <= n; i++ {
		count += strings.Count(strconv.Itoa(i*i), strconv.Itoa(d))
	}
	fmt.Println(count)
}
