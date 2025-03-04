package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var A, B, N int
	fmt.Scan(&A, &B, &N)
	if math.Ceil(float64(A)) > math.Ceil((float64(B) / float64(N))) {
		fmt.Println("Yes")
		os.Exit(0)
	}
	fmt.Println("No")
}
