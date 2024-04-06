package main

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) (res string) {
	d := strings.Split(in, " ")
	max, min := -10000, 10000
	for i := range d {
		c, _ := strconv.Atoi(d[i])
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	res = strconv.Itoa(max) + " " + strconv.Itoa(min)
	return
}

func main() {
	fmt.Println(HighAndLow("-3 -1"))
}
