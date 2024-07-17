package main

import (
	"fmt"
)

func lcg(e, a, m int) int {
	return (a*e + 11) % m
}

func generator(a, m int) <-chan int {
	out := make(chan int)
	go func() {
		seed := 0
		for {
			seed = lcg(seed, a, m)
			result := (abs(seed%3-1)*5 + abs(seed%3)*2) % 8
			out <- result
		}
	}()
	return out
}

func main() {
	var n, k, a, m int
	fmt.Scan(&n, &k, &a, &m)

	coins := generator(a, m)
	coinsReceived := 0
	coinsUsed := 0
	candiesReceived := 0

	for candiesReceived < n {
		coin := <-coins
		coinsUsed++
		coinsReceived += coin

		if coinsReceived >= 3 && coinsReceived >= k*3 {
			candies := coinsReceived / 3
			candiesReceived += candies
			coinsReceived -= candies * 3
		}
	}

	fmt.Println(coinsUsed)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
