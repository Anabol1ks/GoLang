package main

import "fmt"

func Hero(bullets, dragons int) bool {
	return bullets/2 >= dragons
}

func main() {
	fmt.Println(Hero(1500, 751))
}
