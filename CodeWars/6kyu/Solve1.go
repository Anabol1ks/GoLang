package main

import "fmt"

func delDub(data [][]int) [][]int {
	newData := [][]int{}
	for _, sl := range data {
		mNum := make(map[int]bool)
		newArr := []int{}
		for _, num := range sl {
			if !mNum[num] {
				mNum[num] = true
				newArr = append(newArr, num)
			}
		}
		newData = append(newData, newArr)
	}
	return newData
}

func Solve(data [][]int) (res int) {
	res = 1
	newData := delDub(data)
	for _, d := range newData {
		res *= len(d)
	}
	return
}

func main() {
	fmt.Println(Solve([][]int{{1, 2, 3}, {3, 4, 6, 6, 7}, {8, 9, 10, 12, 5, 6}}))
}
