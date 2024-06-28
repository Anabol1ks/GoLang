package main

import (
	"fmt"
	"os"
)

func main() {
	var A, B, N int
	fmt.Scan(&A, &B, &N)
	for i := 1; i <= N; i++ {
		for j := 1; j <= N; j++ {
			if float64(A/j) > float64(B/i) {
				fmt.Println("Yes")
				os.Exit(0)
			}
		}
	}
	fmt.Println("No")
	fmt.Println(9 / 2)
}
