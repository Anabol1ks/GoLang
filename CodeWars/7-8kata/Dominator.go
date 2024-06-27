package main

import "fmt"

func Dominator(a []int) int {
	cout := make(map[int]int)
	max := len(a) / 2
	dom := -1
	for _, i := range a {
		cout[i]++
		if cout[i] > max {
			max = cout[i]
			dom = i
		}
	}
	return dom
}

func main() {
	fmt.Println(Dominator([]int{3, 4, 9, 2, 3, 1, 6, 3}))
}
