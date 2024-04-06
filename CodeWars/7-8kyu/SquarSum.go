package main

import "fmt"

func SquareSum(numbers []int) (sum int) {
	for i := range numbers {
		sum += numbers[i] * numbers[i]
	}
	return
}

func main() {
	numbers := []int{1, 2}
	fmt.Println(SquareSum(numbers))
}
