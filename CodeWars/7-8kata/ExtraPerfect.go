package main

import "fmt"

func ExtraPerfect(n int) (res []int) {
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			res = append(res, i)
		}
	}
	return res
}

func main() {
	fmt.Println(ExtraPerfect(7))
}
