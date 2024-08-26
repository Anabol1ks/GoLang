package main

import (
	"fmt"
	"math"
)

func IterPi(epsilon float64) (int, string) {
	PI := math.Pi
	var res int
	var sum float64

	for {
		term := 1.0 / (2.0*float64(res) + 1.0)
		if res%2 == 0 {
			sum += term
		} else {
			sum -= term
		}

		if math.Abs(4*sum-PI) < epsilon {
			break
		}

		res++
	}

	approxPI := 4 * sum
	return res + 1, fmt.Sprintf("%.10f", approxPI)
}

func main() {
	fmt.Println(IterPi(0.05))
}
