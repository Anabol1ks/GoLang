package main

import "fmt"

func CountSheeps(numbers []bool) (count int) {
	for _, v := range numbers {
		if v {
			count++
		}
	}
	return
}

func main() {
	num := []bool{
		true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true,
	}
	fmt.Println(CountSheeps(num))
}
