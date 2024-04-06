package main

import "fmt"

func High(s string) string {

	sigma := make(map[string]int)
	c := 1
	for ch := 'a'; ch <= 'z'; ch++ {
		sigma[string(ch)] = c
		c++
	}
	sumMax := 0
	sum := 0
	for _, i := range s {
		if string(i) == " " {
			sum = 0
		}
		sum += sigma[string(i)]
		if sum > sumMax {
			sumMax = sum
		}
	}
	fmt.Println(sigma, sumMax)

	return "0"
}

func main() {
	fmt.Println(High("b a c"))
}
