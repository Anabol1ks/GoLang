package main

import (
	"fmt"
	"strconv"
)

func StringToNumber(str string) int {
	n, _ := strconv.Atoi(str)
	return n
}

func main() {
	str := "123"
	fmt.Println(StringToNumber(str))
}
