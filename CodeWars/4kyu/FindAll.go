package main

import "fmt"

func FindAll(sumDig, digs int) []int {
	var results []int
	var generate func(current []int, start, sum, digits int)

	generate = func(current []int, start, sum, digits int) {
		if digits == 0 {
			if sum == 0 {
				number := 0
				for _, digit := range current {
					number = number*10 + digit
				}
				results = append(results, number)
			}
			return
		}

		for i := start; i <= 9; i++ {
			if sum-i >= 0 {
				generate(append(current, i), i, sum-i, digits-1)
			}
		}
	}

	generate([]int{}, 1, sumDig, digs)
	if len(results) > 0 {
		res := []int{len(results), results[0], results[len(results)-1]}
		return res
	}
	return nil
}

func main() {
	fmt.Println(FindAll(10, 3))
}
