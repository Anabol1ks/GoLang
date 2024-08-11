package main

import (
	"fmt"
	"math"
	"strconv"
)

func DigPow(n, p int) int {
	nStr := strconv.Itoa(n)
	sum := 0
	for _, num := range nStr {
		ch, _ := strconv.Atoi(string(num))
		sum += int(math.Pow(float64(ch), float64(p)))
		p++
	}
	if sum%n == 0 {
		return sum / n
	}
	return -1
}

func main() {
	fmt.Println(DigPow(46288, 3))
}
