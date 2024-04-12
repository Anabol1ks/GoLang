package main

import (
	"fmt"
)

func main() {
	var w int
	var m []int
	// fmt.Scan(&n, &t)
	fmt.Scanln(&m)
	fmt.Println("dsa")
	fmt.Scanln(&w)

	// time := 0
	// for i, v := range m {
	// 	t_Up := 0
	// 	for _, o := range m[:w-1] {
	// 		t_Up += int(math.Abs(m[o] - m[o+1]))
	// 	}
	// }
	for i := range m[:w-1] {
		fmt.Println(m[i])
	}
}
