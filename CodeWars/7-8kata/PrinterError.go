package main

import (
	"fmt"
	"strconv"
)

func PrinterError(s string) string {
	var b, size int
	color := make(map[string]bool)
	for i := 'a'; i <= 'm'; i++ {
		color[string(i)] = true
	}
	for _, i := range s {
		size++
		if color[string(i)] == false {
			b++
		}
	}
	res := strconv.Itoa(b) + "/" + strconv.Itoa(size)
	return res
}

func main() {
	fmt.Println(PrinterError("aaaxbbbbyyhwawiwjjjwwm"))
}
