package main

import "fmt"

func Past(h, m, s int) int {
	return (((h*60)+m)*60 + s) * 1000
}

func main() {
	fmt.Println(Past(0, 1, 1))
}
