package main

import (
	"fmt"
	"math"
)

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
			sum += uint64(math.Pow(float64(s%10), float64(l)))
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
