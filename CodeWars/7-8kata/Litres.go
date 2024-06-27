package main

import "fmt"

func Litres(time float64) int {
	return int(time / 2)
}

func main() {
	fmt.Println(Litres(6.7))
}
