package main

import (
	"fmt"
	"strconv"
)

func ConvertFracts(a [][]int) (res string) {
	var count, o int
	for i := 1; ; i++ {
		count = 0
		for n := 0; n < len(a); n++ {
			d := a[n][1]
			if i%d == 0 {
				count++
			}
		}
		if count == len(a) {
			o = i
			break
		}
	}
	for i, _ := range a {
		d := a[i][1]
		del := o / d
		zn := strconv.Itoa(a[i][0] * del)
		res += "(" + zn + "," + strconv.Itoa(o) + ")"
	}
	return
}

func main() {
	fmt.Println(ConvertFracts([][]int{{69, 130}, {87, 1310}, {30, 40}}))
}
