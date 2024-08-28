package main

import (
	"fmt"
)

func powInt(base, exp uint64) uint64 {
	result := uint64(1)
	for exp > 0 {
		result *= base
		exp--
	}
	return result
}

func lennum(n uint64) int {
	count := 0
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}

func SumDigPow(a, b uint64) (res []uint64) {
	for i := a; i <= b; i++ {
		sum := uint64(0)
		s := i
		l := lennum(s)
		for s > 0 {
			sum += powInt(s%10, uint64(l))
			l--
			s = s / 10
		}
		if sum == i {
			res = append(res, i)
		}
	}
	if len(res) == 0 {
		return nil
	}
	return
}

func main() {
	fmt.Println(SumDigPow(12157692622039623308, 12157692622039625693))
}
