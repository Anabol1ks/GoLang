package main

import "fmt"

func HoopCount(n int) string {
	if n >= 10 {
		return "Great, now move on to tricks"
	}
	return "Keep at it until you get it"
}

func main() {
	fmt.Println(HoopCount(0))
}
