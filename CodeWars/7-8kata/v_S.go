package main

import "fmt"

func GetSize(w, h, d int) (res [2]int) {
	res[0] = 2*w*h + 2*w*d + 2*d*h
	res[1] = w * h * d
	return
}

func main() {
	fmt.Println(GetSize(4, 2, 6))
}
