package main

import "fmt"

func Add(n int) func(int) int {
	return func(i int) int {
		return n + i
	}
}

func main() {
	fmt.Println(Add(1)(3)) // 4

}
