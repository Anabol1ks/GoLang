package main

import (
	"fmt"
	"strings"
)

func Points(games []string) (score int) {
	for _, v := range games {
		d := v
		s := strings.Split(d, ":")
		x, y := s[0], s[1]
		if x > y {
			score += 3
		}
		if x == y {
			score++
		}
	}
	return
}

func main() {
	games := []string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}
	score := 0
	for _, v := range games {
		d := v
		s := strings.Split(d, ":")
		x := s[0]
		y := s[1]
		if x > y {
			score += 3
		}
		if x == y {
			score++
		}
	}
	fmt.Println(score)
}
