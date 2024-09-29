package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)
	if 1 <= N && N <= int(2*math.Pow(10, 9)) {
		if N%2 == 0 {
			fmt.Println(N / 2)
		} else {
			fmt.Println(N/2 + 1)
		}
	} else {
		fmt.Println("Error")
	}
}
