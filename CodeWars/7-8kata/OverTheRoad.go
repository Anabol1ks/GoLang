package main

import "fmt"

func OverTheRoad(address int, n int) int {
	d := n * 2
	if address%2 == 0 {
		return d - address + 1
	}
	if address%2 == 1 {
		return d - (address - 1)
	}
	return 0
}

func main() {
	fmt.Println(OverTheRoad(4, 4))
}

// 1 8
// 3 6
// 5 4
// 7 2
