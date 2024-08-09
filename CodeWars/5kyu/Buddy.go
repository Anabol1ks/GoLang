package main

import "fmt"

func proper(n int) (res int) {
	for i := 1; i*i <= n; i++ {
		if n%i == 0 {
			res += i
			if i != 1 && i != n/i {
				res += n / i
			}
		}
	}
	return
}

func Buddy(start, limit int) (res []int) {
	for n := start; n <= limit; n++ {
		m := proper(n) - 1
		if m > n {
			if proper(m) == n+1 {
				res = append(res, n, m)
				return
			}
		}
	}
	return
}

func main() {
	fmt.Println(Buddy(6379, 8275))
}
