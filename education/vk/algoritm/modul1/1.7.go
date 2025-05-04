package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	res := -1
	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			res = arr[i]
		}
	}
	fmt.Println(res)
}
