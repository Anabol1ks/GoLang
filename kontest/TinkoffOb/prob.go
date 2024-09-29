package main

import "fmt"

func main() {
	num1, num2 := 0, 0
	fmt.Scan(&num1, &num2)
	if num1 > -32000 && num2 > -32000 {
		if num1 < 32000 && num2 < 32000 {
			fmt.Println(num1 + num2)
		} else {
			fmt.Println("Error")
		}

	} else {
		fmt.Println("Error")
	}
}
