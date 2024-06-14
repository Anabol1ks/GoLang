package main

import (
	"fmt"
	"math"
)

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	m := 0
	s := 0.00
	saving := float64(savingperMonth)
	priceOld := float64(startPriceOld)
	priceNew := float64(startPriceNew)
	loss := percentLossByMonth
	for {
		m++
		if m%2 == 0 {
			loss += 0.5
		}
		s += saving
		priceOld -= priceOld * loss / 100
		priceNew -= priceNew * loss / 100
		fmt.Println(s, priceOld, priceNew)
		fmt.Println("-")
		fmt.Println(s+priceOld-priceNew, loss)
		fmt.Println("+")

		if s+priceOld >= priceNew {
			break
		}
	}
	res := s + priceOld - priceNew

	return [2]int{m, int(math.Round(res))}
}

func main() {
	fmt.Println(NbMonths(7500, 32000, 300, 1.55))
}
