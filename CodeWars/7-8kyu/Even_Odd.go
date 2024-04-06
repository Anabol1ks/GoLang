package main

import "fmt"

func EvenOrOdd(number int) string {
	str := ""
	if number%2 == 0 {
		str = "Even"
	} else {
		str = "Odd"
	}
	return str
}
func main() {
	s := EvenOrOdd(5)
	fmt.Println(s)
}
