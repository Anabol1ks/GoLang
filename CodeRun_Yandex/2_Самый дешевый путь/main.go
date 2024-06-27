package main

import "fmt"

type State struct {
	n, m int
}

func main() {
	var n, m int
	fmt.Scanln(&n, &m)
	if (1 > n || n > 20) || (1 > m || m > 20) {
		panic("Размеры таблицы должны быть 1<= N, M<= 20")
	}
	table := make([][]int, n)
	for i := range table {
		table[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&table[i][j])
		}
	}
	for i := range len(table) {
		fmt.Println(table[i])
	}
}
