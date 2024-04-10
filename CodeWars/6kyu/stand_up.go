package main

import "fmt"

func stand_up(kol int, h []int, new int, name string) int {
	if kol != len(h) {
		return 0
	}
	fmt.Println(name)
	sum := 0
	for i, s := range h {
		sum += (i + 1) * s
	}
	n_sum, s_max := 0, 0
	for i := 0; i <= len(h); i++ {
		in := i
		farr := make([]int, in)
		copy(farr, h[:in])
		sarr := make([]int, len(h)-in)
		copy(sarr, h[in:])
		newH := append(farr, new)
		newH = append(newH, sarr...)
		for j, s := range newH {
			n_sum += (j + 1) * s
		}
		if n_sum > s_max {
			s_max = n_sum
		}
		n_sum = 0

	}
	return s_max - sum
}

func main() {

	fmt.Println(stand_up(3, []int{1, 4, 3}, 2, "Гурам"))
}
