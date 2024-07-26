package main

import (
	"bufio"
	"fmt"
	"math"
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
		var n, k, m int
		fmt.Fscan(in, &n, &k)
		fmt.Fscan(in, &m)
		arr := make([]int, m)
		for j := 0; j < m; j++ {
			var ch int
			fmt.Fscan(in, &ch)
			arr[j] = int(math.Pow(2, float64(ch)))
		}

		// Сортировка коробок по убыванию веса
		sort.Sort(sort.Reverse(sort.IntSlice(arr)))

		res := 0
		for len(arr) > 0 {
			machines := make([]int, n) // Текущие машины с их остаточной грузоподъемностью
			for j := 0; j < len(arr); {
				loaded := false
				for m := 0; m < n; m++ {
					if machines[m]+arr[j] <= k {
						machines[m] += arr[j]
						arr = append(arr[:j], arr[j+1:]...) // Удаляем коробку из списка
						loaded = true
						break
					}
				}
				if !loaded {
					j++
				}
			}
			res++
		}
		fmt.Fprintln(out, res)
	}
}
