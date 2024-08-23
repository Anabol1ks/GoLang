package main

import (
	"fmt"
	"math"
)

func Movie(card, ticket int, perc float64) (res int) {
	var a int
	b := float64(card)
	predTicket := float64(ticket)
	for math.Ceil(b) >= float64(a) {
		res++
		a += ticket
		b += predTicket * perc
		predTicket = predTicket * perc
	}
	return
}

func main() {
	fmt.Println(Movie(500, 15, 0.9))
}
