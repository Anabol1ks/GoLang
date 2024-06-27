package main

import "fmt"

func QuarterOf(month int) int {
	quart := 0
	switch {
	case 1 > month || month > 12:
		fmt.Println("Ошибка")
	case 3 >= month:
		quart = 1
	case 6 >= month:
		quart = 2
	case 9 >= month:
		quart = 3
	case 12 >= month:
		quart = 4
	}
	return quart
}

func main() {
	fmt.Println(QuarterOf(4))
}
