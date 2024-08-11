package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func SumOfIntegersInString(strng string) (sum int) {
	re := regexp.MustCompile(`\d+`)
	nums := re.FindAllString(strng, -1)

	for _, num := range nums {
		num, _ := strconv.Atoi(num)
		sum += num
	}

	return sum
}

func main() {
	fmt.Println(SumOfIntegersInString("The Great Depression lasted from 1929 to 1939."))
}
