package main

import (
	"fmt"
	"sort"
)

func SortNumbers(numbers []int) []int {
	if numbers == nil {
		return nil
	}
	sort.Ints(numbers)
	return numbers
}

func main() {
	c := []int{1, 2, 4, 3, 6, 5}
	fmt.Println(SortNumbers(c))
}
