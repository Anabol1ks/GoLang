package main

import "fmt"

func FindUniq(arr []float32) float32 {
	for i := range arr {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] != arr[j] {
				fmt.Println("")
				return arr[j]
			}
		}
	}
	return arr[0]
}

func main() {
	fmt.Println(FindUniq([]float32{2.0, 1.0, 1.0, 1.0, 1.0, 1.0}))
}
