package main

import "fmt"

func LoveFunc(flower1, flower2 int) bool {
	return flower1%2 == 1 && flower2%2 == 0 || flower1%2 == 0 && flower2%2 == 1
}

func main() {
	fmt.Println(LoveFunc(6, 12))
}
