package main

import (
	"fmt"
	"os"
)

func main() {
	var n, res int
	fmt.Scan(&n)
	m := make([]int, n)
	l := len(m)
	for i := 0; i < l; i++ {
		fmt.Scan(&m[i])
		if 2 > m[i] || m[i] > 5 {
			fmt.Println("Ошибка")
			os.Exit(0)
		}
	}
	if l <= 7 {
		for i := range m {
			if m[i] <= 3 {
				res = -1
				break
			}
			if m[i] == 5 {
				res++
			}
		}
	}
	count := 0
	if l > 7 {
		for i := 0; i < l-6; i++ {
			for j := i; j < i+7; j++ {
				if m[j] == 5 {
					count++
				}
				if count > res {
					res = count
				}
				if m[j] <= 3 {
					res = -1
					break
				}
			}

			count = 0
		}
	}
	fmt.Println(res)

}
