package main

import "fmt"

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
	i := 5
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

func Gap(g, m, n int) []int {
	lastPrime := -1
	for i := m; i <= n; i++ {
		if isPrime(i) {
			if lastPrime != -1 && i-lastPrime == g {
				return []int{lastPrime, i}
			}
			lastPrime = i
		}
	}
	return nil
}

func main() {
	fmt.Println(Gap(4, 100, 110))
}
