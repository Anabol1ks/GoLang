package main

import "fmt"

func Gps(s int, x []float64) (res int) {
	var do_res float64
	for i := 1; i < len(x); i++ {
		f := x[i] - x[i-1]
		do_res = (3600 * f) / float64(s)
		if int(do_res) > res {
			res = int(do_res)
		}
	}
	return
}

func main() {
	var x = []float64{0.0, 0.19, 0.5, 0.75, 1.0, 1.25, 1.5, 1.75, 2.0, 2.25}
	fmt.Println(Gps(15, x))
}
