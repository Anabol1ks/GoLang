package main

import "fmt"

func CountPositivesSumNegatives(numbers []int) []int {
	var res []int
	c, s := 0, 0
	for _, n := range numbers {
		if n > 0 {
			c += 1
		} else {
			s += n
		}
	}
	res = append(res, c, s)
	return res
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	fmt.Println(CountPositivesSumNegatives(numbers))
}
