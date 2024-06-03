package main

import (
	"fmt"
	"strconv"
)

func countSheep(num int) (res string) {
	for i := 1; i <= num; i++ {
		res += strconv.Itoa(i) + " sheep..."
	}
	return
}

func main() {
	fmt.Println(countSheep(0))
}
