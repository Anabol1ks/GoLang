package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var t int
	fmt.Fscan(in, &t)

	for i := 0; i < t; i++ {
		var n int
		res := "yes"

		_, err := fmt.Fscan(in, &n)
		if err != nil {
			fmt.Fprintln(out, "no")
			in.ReadString('\n')
			continue
		}

		arr1 := make([]int, n)
		arr2 := make([]int, n)

		// Чтение элементов для arr1
		for j := 0; j < n; j++ {
			var num int
			_, err := fmt.Fscan(in, &num)
			if err != nil {
				res = "no"
				in.ReadString('\n')
				break
			}
			arr1[j] = num
		}

		// Чтение элементов для arr2
		if res == "yes" {
			for j := 0; j < n; j++ {
				var num int
				_, err := fmt.Fscan(in, &num)
				if err != nil {
					res = "no"
					in.ReadString('\n')
					break
				}
				arr2[j] = num
			}
		}

		if res == "yes" {
			sort.Ints(arr1)
			for j := 0; j < n; j++ {
				if arr1[j] != arr2[j] {
					res = "no"
					break
				}
			}
		}

		fmt.Fprintln(out, res)
	}
}
