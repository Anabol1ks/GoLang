package main

import "fmt"

func SequenceSum(start, end, step int) (res int) {
	for i := start; i <= end; i += step {
		res += i
	}
	return
}

func main() {
	fmt.Println(SequenceSum(2, 6, 2))
}
