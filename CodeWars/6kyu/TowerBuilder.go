package main

import (
	"fmt"
	"strings"
)

func TowerBuilder(nFloors int) (tower []string) {
	for i := 1; i <= nFloors; i++ {
		space := strings.Repeat(" ", nFloors-i)
		z := strings.Repeat("*", 2*i-1)
		tower = append(tower, space+z+space)
	}
	return
}

func main() {
	fmt.Println(TowerBuilder(4))
}
