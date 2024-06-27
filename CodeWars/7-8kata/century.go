package main

import "fmt"

func century(year int) int {
	if year%100 >= 1 {
		return year/100 + 1
	}
	return year / 100
}

func main() {
	fmt.Println(century(2000))
}
