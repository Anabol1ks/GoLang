package main

import (
	"fmt"
	"strconv"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func reduceFraction(numerator, denominator int) (int, int) {
	gcdValue := gcd(numerator, denominator)
	return numerator / gcdValue, denominator / gcdValue
}

func ConvertFracts(a [][]int) (res string) {
	for i := range a {
		a[i][0], a[i][1] = reduceFraction(a[i][0], a[i][1])
	}

	var count, lcm int
	for i := 1; ; i++ {
		count = 0
		for _, frac := range a {
			d := frac[1]
			if i%d == 0 {
				count++
			}
		}
		if count == len(a) {
			lcm = i
			break
		}
	}

	for _, frac := range a {
		num := frac[0]
		denom := frac[1]
		multiplier := lcm / denom
		newNum := num * multiplier
		res += "(" + strconv.Itoa(newNum) + "," + strconv.Itoa(lcm) + ")"
	}
	return
}

func main() {
	fmt.Println(ConvertFracts([][]int{{69, 130}, {87, 1310}, {30, 40}}))
}
