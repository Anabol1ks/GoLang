package main

import (
	"fmt"
	"strconv"
)

func MultiTable(number int) (table string) {
	for i := 1; i <= 10; i++ {
		table += strconv.Itoa(i) + " * " + strconv.Itoa(number) + " = " + strconv.Itoa(i*number)
		if i != 10 {
			table += "\n"
		}
	}
	return
}

func main() {
	fmt.Println(MultiTable(5))
}
