package main

import "fmt"

func main() {
	var w, b int
	fmt.Scan(&b, &w)
	cv := w + b
	for n := 1; n < cv/2+1; n++ {
		if cv%n == 0 {
			m := cv / n
			if 2*n+2*m-4 == b && (n-2)*(m-2) == w {
				fmt.Println(m, n)
				return
			}
			if 2*m+2*n-4 == b && (m-2)*(n-2) == w {
				fmt.Println(m, n)
				return
			}
		}
	}
}
