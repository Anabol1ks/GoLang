package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0
	length := len(flowerbed)
	for i := 0; i < length; i++ {
		if length == 1 {
			if flowerbed[i] == 0 {
				count++
			}
			break
		}
		if i == 0 {
			if flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				fmt.Println(flowerbed, "1")
				flowerbed[i] = 1
				count++
			}
		} else if i == length-1 {
			if flowerbed[i] == 0 && flowerbed[i-1] == 0 {
				flowerbed[i] = 1
				count++
			}
		} else {
			if flowerbed[i-1] == 0 && flowerbed[i] == 0 && flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				count++
			}
		}
	}

	if count < n {
		return false
	}
	return true
}

func main() {
	fmt.Println(canPlaceFlowers([]int{0, 0, 1, 0, 0}, 1))
}
