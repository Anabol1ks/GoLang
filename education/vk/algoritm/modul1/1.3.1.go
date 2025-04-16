package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	noZero := make([]int, 0)
	for _, i := range arr {
		if i != 0 {
			noZero = append(noZero, i)
		}
	}
	zero := n - len(noZero)
	for i := 0; i < zero; i++ {
		noZero = append(noZero, 0)
	}
	var sb strings.Builder
	for i, num := range noZero {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(strconv.Itoa(num))
	}
	fmt.Println(sb.String())
}
