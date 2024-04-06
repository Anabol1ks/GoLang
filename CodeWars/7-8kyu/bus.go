package main

import "fmt"

func Number(busStops [][2]int) int {
	sum, min := 0, 0
	for i := range busStops {
		sum += busStops[i][0]
		min += busStops[i][1]
	}
	return sum - min
}

func main() {
	d := [][2]int{
		{10, 0},
		{3, 5},
		{5, 8},
	}
	fmt.Println(Number(d))

}
