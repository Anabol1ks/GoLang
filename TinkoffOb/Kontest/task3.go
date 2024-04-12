package main

import (
	"fmt"
	"math"
)

func main() {
	var n, t, w int
	fmt.Scan(&n, &t)
	m := make([]int, n)
	for i := 0; i < len(m); i++ {
		fmt.Scan(&m[i])
	}
	fmt.Scan(&w)

	fmt.Println(m)
	time := 0
	t_Up := 0
	for j := range m[:w-1] {
		sub := m[j] - m[j+1]
		t_Up += int(math.Abs(float64(sub)))
	}
	if t_Up >= t {
		m = append(m[:w], m[w+1:]...)
	}
	fmt.Println(m)
	for i := 0; i < len(m); i++ {
		if i != len(m)-1 {
			sub := m[i] - m[i+1]
			time += int(math.Abs(float64(sub)))
		}
	}
	fmt.Println(time)

}
