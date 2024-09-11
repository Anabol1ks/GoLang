package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}

func countDel(num int) int {
	count := 0
	sqrt := int(math.Sqrt(float64(num)))
	for i := 1; i <= sqrt; i++ {
		if num%i == 0 {
			if i*i == num {
				count++
			} else {
				count += 2
			}
		}
	}
	return count
}

func ResCount(l, r int) int {
	count := 0
	for num := l; num <= r; num++ {
		if num < 2 {
			continue
		}
		if isPrime(num) {
			continue
		}
		divCount := countDel(num)
		if isPrime(divCount) {
			count++
		}
	}
	return count
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var l, r int
	fmt.Fscan(in, &l, &r)
	fmt.Fprint(out, ResCount(l, r))
}
