package main

import "fmt"

func main() {
	var n, sum int
	fmt.Scan(&n)
	mas := make([]int, n)
	for i := 0; i < len(mas); i++ {
		fmt.Scan(&mas[i])
	}
	for i := range mas {
		sum += i
	}
	if sum%2 == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
