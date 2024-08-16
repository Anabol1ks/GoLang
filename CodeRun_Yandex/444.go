package main

import (
	"bufio"
	"fmt"
	"os"
)

func isUniform(slice []string) bool {
	first := slice[0]
	for _, val := range slice {
		if val != first {
			return false
		}
	}
	return true
}

func Row(pole [][]string, n, m int) bool {
	for i := 0; i < n; i++ {
		for j := 0; j <= m-5; j++ {
			if isUniform(pole[i][j : j+5]) {
				return true
			}
		}
	}
	return false
}

func checkColumn(pole [][]string, n, m int) bool {
	for j := 0; j < m; j++ {
		for i := 0; i <= n-5; i++ {
			column := []string{}
			for k := 0; k < 5; k++ {
				column = append(column, pole[i+k][j])
			}
			if isUniform(column) {
				return true
			}
		}
	}
	return false
}

func checkAllDiagonals(pole [][]string, n, m int) bool {
	for startRow := 0; startRow <= n-5; startRow++ {
		for startCol := 0; startCol <= m-5; startCol++ {
			if checkDiagonal(pole, startRow, startCol, 1, 1) {
				return true
			}
		}
	}

	for startRow := 4; startRow < n; startRow++ {
		for startCol := 0; startCol <= m-5; startCol++ {
			if checkDiagonal(pole, startRow, startCol, -1, 1) {
				return true
			}
		}
	}

	return false
}

func checkDiagonal(pole [][]string, startRow, startCol, rowStep, colStep int) bool {
	diagonal := []string{}
	for i := 0; i < 5; i++ {
		diagonal = append(diagonal, pole[startRow+i*rowStep][startCol+i*colStep])
	}
	return isUniform(diagonal)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	res := "No"
	fmt.Fscan(in, &n)
	fmt.Fscan(in, &m)

	pole := make([][]string, n)
	for i := range pole {
		pole[i] = make([]string, m)
	}

	for i := 0; i < n; i++ {
		var str string
		fmt.Fscan(in, &str)
		if len(str) == m {
			for j, ch := range str {
				pole[i][j] = string(ch)
			}
		}
	}
	if Row(pole, n, m) {
		res = "Yes"
	}
	if checkColumn(pole, n, m) {
		res = "Yes"
	}
	if checkAllDiagonals(pole, n, m) {
		res = "Yes"
	}

	fmt.Fprintln(out, res)
}
