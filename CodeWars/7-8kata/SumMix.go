package main

import (
	"fmt"
	"strconv"
)

type any interface{}

func SumMix(arr []any) (sum int) {
	for _, i := range arr {
		switch v := i.(type) {
		case int:
			sum += v
		case string:
			num, _ := strconv.Atoi(v)
			sum += num
		}
	}
	return
}

func main() {
	fmt.Println(SumMix([]any{"3", 6, 6, 0, "5", 8, 5, "6", 2, "0"}))
}
