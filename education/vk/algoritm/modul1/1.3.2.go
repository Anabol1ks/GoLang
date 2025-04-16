package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n, element int
	fmt.Scan(&n)

	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Scan(&element)

	for i := 0; i < n; i++ {
		if arr[i] == element {
			arr = append(arr[:i], arr[i+1:]...)
			i = -1
			n--
		}
	}

	var sb strings.Builder
	for i, num := range arr {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(strconv.Itoa(num))
	}
	fmt.Println(sb.String())
}
