package main

import "fmt"

func PositiveSum(numbers []int) (sum int) {
	for _, n := range numbers {
		if n > 0 {
			sum += n
		}
	}
	return
}

func main() {
	numbers := []int{1, -4, 7, -12}
	fmt.Println(PositiveSum(numbers))
}
