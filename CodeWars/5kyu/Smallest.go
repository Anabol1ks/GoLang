package main

import (
	"fmt"
	"math"
	"strconv"
)

func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func insert(slice []string, index int, value string) []string {
	return append(slice[:index], append([]string{value}, slice[index:]...)...)
}

func Smallest(n int64) (res []int64) {
	i := n
	arr := []string{}
	for i > 0 {
		ch := strconv.Itoa(int(i % 10))
		arr = append([]string{ch}, arr...)
		i = i / 10
	}
	str := ""
	min := math.MaxInt
	var in1, in2, num int
	for i, n := range arr {
		newSlice := make([]string, len(arr))
		copy(newSlice, arr)

		newSlice = remove(newSlice, i)
		for j := range newSlice {
			tempSlice := make([]string, len(newSlice))
			copy(tempSlice, newSlice)
			if j >= i {
				j++
			}
			tempSlice = insert(tempSlice, j, n)
			for _, s := range tempSlice {
				str += s
			}
			num, _ = strconv.Atoi(str)
			str = ""
			if num < min {
				min = num
				in1, in2 = i, j
			}
		}
	}
	res = append(res, int64(min), int64(in1), int64(in2))
	return
}

func main() {
	fmt.Println(Smallest(261235))
}
