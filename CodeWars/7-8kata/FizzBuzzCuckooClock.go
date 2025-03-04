package main

import (
	"fmt"
	"strconv"
	"strings"
)

func FizzBuzzCuckooClock(time string) string {
	h_m := strings.Split(time, ":")
	h, _ := strconv.Atoi(h_m[0])
	m, _ := strconv.Atoi(h_m[1])
	if h > 12 {
		h = h - 12
	} else if h == 0 {
		h = 12
	}
	if m%15 == 0 && m != 30 && m != 00 {
		return "Fizz Buzz"
	} else if m%3 == 0 && m%5 != 0 {
		return "Fizz"
	} else if m%5 == 0 && m%3 != 0 {
		return "Buzz"
	} else if m == 0 {
		res := strings.Repeat("Cuckoo ", h)
		return res[:len(res)-1]
	} else if m == 30 {
		return "Cuckoo"
	}
	return "tick"
}

func main() {
	fmt.Println(FizzBuzzCuckooClock("21:00"))
}
