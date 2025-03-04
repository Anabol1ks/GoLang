package main

import (
	"bufio"
	"fmt"
	"os"
)

func isUniform(slice []string) bool {
	first := slice[0]
	if first == "." {
		return false
	}
	for _, val := range slice {
		if val != first {
			return false
		}
	}
	return true
}

func checkRow(pole [][]string, n, m int) bool {
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

func checkDiagonal(pole [][]string, n, m int) bool {
	for startRow := 0; startRow <= n-5; startRow++ {
		for startCol := 0; startCol <= m-5; startCol++ {
			diagonal1 := true
			diagonal2 := true
			for k := 0; k < 5; k++ {
				if pole[startRow+k][startCol+k] != pole[startRow][startCol] || pole[startRow][startCol] == "." {
					diagonal1 = false
				}
				if pole[startRow+k][startCol+4-k] != pole[startRow][startCol+4] || pole[startRow][startCol+4] == "." {
					diagonal2 = false
				}
			}
			if diagonal1 || diagonal2 {
				return true
			}
		}
	}
	return false
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n, m int
	fmt.Fscan(in, &n, &m)

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

	if checkRow(pole, n, m) || checkColumn(pole, n, m) || checkDiagonal(pole, n, m) {
		fmt.Fprintln(out, "Yes")
	} else {
		fmt.Fprintln(out, "No")
	}
}
