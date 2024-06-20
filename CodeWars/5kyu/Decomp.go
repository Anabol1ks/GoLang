package main

import (
	"fmt"
	"strconv"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sieveOfEratosthenes(n int) []int {
	primes := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		primes[i] = true
	}

	for p := 2; p*p <= n; p++ {
		if primes[p] {
			for i := p * p; i <= n; i += p {
				primes[i] = false
			}
		}
	}

	var result []int
	for p := 2; p <= n; p++ {
		if primes[p] {
			result = append(result, p)
		}
	}
	return result
}

func primeFactorCount(n, p int) int {
	count := 0
	power := p
	for power <= n {
		count += n / power
		power *= p
	}
	return count
}

func Decomp(n int) string {
	primes := sieveOfEratosthenes(n)
	factorization := make(map[int]int)

	for _, p := range primes {
		count := primeFactorCount(n, p)
		if count > 0 {
			factorization[p] = count
		}
	}

	res := ""
	for i, p := range primes {
		if i != 0 {
			res += " * "
		}
		if factorization[p] > 1 {
			res += strconv.Itoa(p) + "^" + strconv.Itoa(factorization[p])
		} else {
			res += strconv.Itoa(p)
		}
	}
	return res
}

func main() {
	fmt.Println(Decomp(12))
}
