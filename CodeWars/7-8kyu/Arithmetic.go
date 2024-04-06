package main

import "fmt"

func Arithmetic(a int, b int, operator string) int {
	switch operator {
	case "add":
		return a + b
	case "subtract":
		return a - b
	case "multiply":
		return a * b
	case "divide":
		return a / b
	default:
		return 0
	}
}

func main() {
	fmt.Println(Arithmetic(5, 2, "multiply"))
}
