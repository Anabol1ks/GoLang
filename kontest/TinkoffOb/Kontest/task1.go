package main

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Scanln(&A, &B, &C, &D)
	if 1 <= A && A <= 100 && 1 <= B && B <= 100 && 1 <= C && C <= 100 && 1 <= D && D <= 100 {
		if D > B {
			A += (D - B) * C
		}
		fmt.Println(A)
	} else {
		fmt.Println("Ошибка")
	}
}
