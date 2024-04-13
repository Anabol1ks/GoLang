package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	matrix := make([][]int, n)
	for i := range matrix {
		matrix[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j])
		}
	}

	matrix2 := make([][]int, m)
	for i := range matrix2 {
		matrix2[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			matrix2[i][j] = matrix[n-1-j][i]
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", matrix2[i][j])
		}
		fmt.Println()
	}
}
