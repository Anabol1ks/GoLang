package main

import (
	"fmt"
	"strconv"
)

func main() {

	d := "1234"
	_, err := strconv.Atoi(d)
	if err != nil {
		fmt.Println("Ошибка")
	}
	fmt.Println(err)
}
