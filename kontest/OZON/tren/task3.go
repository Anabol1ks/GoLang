package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var l, ver, res int
		kol := make(map[int][]int)
		fmt.Fscan(in, &l)
		arr := make([]int, l)
		for j := 0; j < l; j++ {
			fmt.Fscan(in, &arr[j])
		}
		status := make(map[int]string)
		status[1] = "ver"
		for n := 0; n < len(arr); n++ {
			switch status[1] {
			case "ver":
				ver = arr[n]
				status[1] = "kDet"
			case "kDet":
				det := make([]int, arr[n])
				c := 0
				for d := 0; d < arr[n]; d++ {
					det[d] = arr[n+d+1]
					c++
				}
				kol[ver] = det
				status[1] = "ver"
				n += c
				c = 0
			}

		}
		for m, _ := range kol {
			c := 0
			for _, m1 := range kol {
				for _, n := range m1 {
					if m == n {
						c++
					}
				}
			}
			if c == 0 {
				res = m
			}
			c = 0
		}
		fmt.Fprintln(out, res)
	}
}

// 3 0 1 0 5 2 2 6 4 3 5 1 3 2 0 6 0
