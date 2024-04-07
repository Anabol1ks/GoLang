package main

import "fmt"

func Tribonacci(signature [3]float64, n int) []float64 {
	if n == 0 {
		return []float64{}
	} else if n <= 3 {
		return signature[:n]
	}

	res := []float64{}
	res = append(res, signature[0], signature[1], signature[2])
	for i := 0; i < n-3; i++ {
		res = append(res, (res[i] + res[i+1] + res[i+2]))
	}
	return res
}

func main() {
	fmt.Println(Tribonacci([3]float64{3, 2, 1}, 10))
}
