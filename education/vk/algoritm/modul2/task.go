package main

import "fmt"

func main() {
	// arr := []int{1, 3, 5, 7, 8, 9, 10, 13, 14, 15, 22, 24, 24, 24, 24, 24, 24, 27, 33}
	// left := 0
	// right := len(arr) - 1
	// fmt.Println(binarySearch(arr, left, right, 24))
	// fmt.Println(binarySearchRight(arr, 24))
	// fmt.Println(binarySearchLeft(arr, 24))

	// 2.2 видос
	// Найти ближайший корень
	// fmt.Println(binarySearchSqrt(66))

	// задача про дипломы
	// fmt.Println(binarySearchDip(3, 4, 9))

	// Очень легкая задача
	// fmt.Println(copyTime(12, 2, 4))

	// 2.3
	// тернарный поиск
	// data := []int{1, 3, 4, 5, 7, 9, 11, 15, 17, 20, 22}
	// fmt.Println(ternarySearch(data, 0, len(data)-1, 22))

	// // экспоненциальный поиск
	// data := []int{1, 3, 4, 5, 7, 9, 11, 15, 17, 20, 22}
	// fmt.Println(binarySearchExp(data, 11))

	// 2.4 сортировки
	arr := []int{4, 1, 41, 3, 5, 21, 2, 9, 33}
	fmt.Println(selectionSort(arr))
	fmt.Println(insertionSort(arr))
	fmt.Println(bubbleSort(arr))
	fmt.Println(shellSort(arr))
}

func shellSort(arr []int) []int {
	gap := len(arr) / 2
	for gap > 0 {
		for i := gap; i < len(arr); i++ {
			m_gap := i
			for m_gap >= gap && arr[m_gap] < arr[m_gap-gap] {
				arr[m_gap], arr[m_gap-gap] = arr[m_gap-gap], arr[m_gap]
				m_gap -= gap
			}
		}
		gap = gap / 2
	}
	return arr
}

func bubbleSort(arr []int) []int {
	var sorted bool
	for !sorted {
		sorted = true
		for i := 0; i < len(arr)-1; i++ {
			if arr[i] > arr[i+1] {
				arr[i], arr[i+1] = arr[i+1], arr[i]
				sorted = false
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {
	len := len(arr)
	for i := 0; i < len-1; i++ {
		min := i
		for j := i + 1; j < len; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
			j--
		}
	}
	return arr
}

// func exponentialSearch(data []int, target int) (int, int) {
// 	border := 1
// 	for border < len(data)-1 {
// 		if data[border] < target {
// 			border *= 2
// 		} else if data[border] >= target {
// 			return border / 2, border
// 		}
// 	}
// 	return 0, 0
// }

// func binarySearchExp(arr []int, target int) int {
// 	left, right := exponentialSearch(arr, target)
// 	if left > right {
// 		return -1
// 	}
// 	middle := (left + right) / 2
// 	if arr[middle] == target {
// 		return middle
// 	}
// 	if arr[middle] > target {
// 		return binarySearch(arr, left, middle-1, target)
// 	} else {
// 		return binarySearch(arr, middle+1, right, target)
// 	}
// }

// func ternarySearch(data []int, left, right, target int) int {
// 	if right >= left {
// 		m1 := left + (right-left)/3
// 		m2 := right - (right-left)/3

// 		if data[m1] == target {
// 			return m1
// 		}
// 		if data[m2] == target {
// 			return m2
// 		}

// 		if target < data[m1] {
// 			return ternarySearch(data, left, m1-1, target)
// 		} else if target > data[m2] {
// 			return ternarySearch(data, m2+1, right, target)
// 		} else {
// 			return ternarySearch(data, m1+1, m2-1, target)
// 		}
// 	}
// 	return -1
// }

// // Очень легкая задача
// func copyTime(n, x, y int) int {
// 	left := 0
// 	right := (n - 1) * max(x, y)
// 	for left+1 < right {
// 		middle := (right + left) / 2
// 		if (middle/x + middle/y) < n-1 {
// 			left = middle
// 		} else {
// 			right = middle
// 		}
// 	}
// 	return right + min(x, y)
// }

// задача про дипломы
// func binarySearchDip(w, h, n int) int {
// 	left := max(w, h)
// 	right := left * n
// 	for left < right {
// 		fmt.Println("dds")
// 		middle := (right + left) / 2
// 		res := (middle / w) * (middle / h)
// 		if res < n {
// 			left = middle
// 		} else {
// 			right = middle
// 		}
// 	}
// 	return right
// }

// Найти ближайший корень
// func binarySearchSqrt(target int) int {
// 	left := 0
// 	right := target
// 	for left <= right {
// 		middle := (left + right) / 2
// 		if middle*middle > target {
// 			right = middle - 1
// 			continue
// 		}
// 		if middle*middle < target {
// 			left = middle + 1
// 			continue
// 		}
// 		return middle
// 	}
// 	return right
// }

// func binarySearch(arr []int, left, right, target int) int {
// 	if left > right {
// 		return -1
// 	}
// 	middle := (left + right) / 2
// 	if arr[middle] == target {
// 		return middle
// 	}
// 	if arr[middle] > target {
// 		return binarySearch(arr, left, middle-1, target)
// 	} else {
// 		return binarySearch(arr, middle+1, right, target)
// 	}
// }

// func binarySearchRight(arr []int, target int) int {
// 	left := 0
// 	right := len(arr) - 1

// 	for left+1 < right {
// 		middle := (left + right) / 2
// 		if arr[middle] <= target {
// 			left = middle
// 		} else {
// 			right = middle
// 		}
// 	}
// 	if arr[right] == target {
// 		return right
// 	}
// 	if arr[left] == target {
// 		return left
// 	}
// 	return -1
// }

// func binarySearchLeft(arr []int, target int) int {
// 	left := 0
// 	right := len(arr) - 1

// 	for left+1 < right {
// 		middle := (left + right) / 2
// 		if arr[middle] < target {
// 			left = middle
// 		} else {
// 			right = middle
// 		}
// 	}
// 	if arr[left] == target {
// 		return left
// 	}
// 	if arr[right] == target {
// 		return right
// 	}
// 	return -1
// }
