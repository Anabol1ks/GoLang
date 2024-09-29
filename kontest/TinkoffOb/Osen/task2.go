package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int
	fmt.Fscan(in, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(in, &arr[i])
	}
	arr2 := make([]int, n)

	if arr[0] == -1 {
		arr2[0] = 1
	} else {
		arr2[0] = arr[0]
	}

	for i := 1; i < n; i++ {
		if arr[i] == -1 {
			if arr[i-1] != -1 && i < n-1 && arr[i+1] != -1 {
				arr2[i] = (arr[i-1] + arr[i+1]) / 2
			} else if arr[i-1] != -1 {
				arr2[i] = arr[i-1] + 1
			} else if i < n-1 && arr[i+1] != -1 {
				arr2[i] = arr[i+1] - 1
			} else {
				arr2[i] = 1
			}
		} else {
			arr2[i] = arr[i]
		}
	}

	res := make([]int, n)
	res[0] = arr2[0]
	for i := 1; i < n; i++ {
		if arr2[i]-arr2[i-1] > res[i-1] {
			res[i] = arr2[i] - arr2[i-1]
		} else {
			res[i] = res[i-1] + 1
		}
	}

	expected := true
	for i := 0; i < n; i++ {
		if res[i] != i+1 {
			expected = false
			break
		}
	}

	if expected {
		fmt.Fprintln(out, "YES")
		for _, v := range res {
			fmt.Fprint(out, v, " ")
		}
		fmt.Fprintln(out)
	} else {
		fmt.Fprintln(out, "NO")
	}
}
