package main

import (
	"fmt"
)

func FakeBin(x string) string {
	new := ""
	f := "5"
	for i := 0; i < len(x); i++ {
		if string(x[i]) < f {
			new += "0"
		} else if string(x[i]) >= f {
			new += "1"
		}
	}
	return new
}

func main() {
	x := "45385593107843568"
	fmt.Println(FakeBin(x))
}
