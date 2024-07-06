package main

import (
	"fmt"
	"math"
)

func TankVol(h, d, vt int) int {
	r := float64(d / 2)
	l := float64(vt) / (3.14 * math.Pow(r, 2))
	fmt.Println(l)
	H := float64(h)
	fmt.Println(math.Acos((r - H) / r))
	V := l * (math.Pow(r, 2)*(math.Acos((r-H)/r)) - (r-H)*math.Sqrt(2*r*H-math.Pow(H, 2)))
	return int(V)
}

func main() {
	fmt.Println(TankVol(40, 120, 3500))
}
