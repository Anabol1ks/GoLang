package main

import "fmt"

func Angle(n int) int {
	return (n - 2) * 180
}

func main() {
	fmt.Println(Angle(4))
}
