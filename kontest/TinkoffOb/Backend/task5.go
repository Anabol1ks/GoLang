package main

import (
	"fmt"
)

func main() {
	var n, m, c int
	res := ""
	fmt.Scanln(&n, &m)
	for i := 0; i < m; i++ {
		var a, b int
		if a <= n && b <= n {
			fmt.Scanln(&a, &b)
			if a < b {
				res = "Yes"
			} else {
				c++
			}
		}
	}
	if c > 0 {
		res = "No"
		fmt.Println(res)
	} else {
		fmt.Println(res)
		num := ""
		for i := 1; i <= n; i++ {
			num += fmt.Sprintf("%d ", i)
		}
		num = num[:len(num)-1]
		fmt.Println(num)
	}

}
