package main

import "fmt"

func Between(a, b int) []int {
	mas := []int{}
	for i := a; i <= b; i++ {
		mas = append(mas, i)
	}
	return mas
}

func main() {
	fmt.Println(Between(1, 5))
}
