package main

func main() {
	// twoSum
	// arr := []int{2, 4, 8, 9, 11, 12, 16, 21}
	// sum := 33
	// fmt.Println(twoSum(arr, sum))

	// // reverseArray
	// arr := []int{2, 4, 8, 9, 11, 12, 16, 21}
	// fmt.Println(reverseArray(arr))

	// очереди, стек, дек

	// 1.7 задача 1
	// isSubsequence
}

func reverseArray(arr []int) []int {
	left, right := 0, len(arr)-1
	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}

func twoSum(arr []int, sum int) (int, int) {
	left := 0
	right := len(arr) - 1
	for left != right {
		tmp := arr[left] + arr[right]
		if tmp == sum {
			return arr[left], arr[right]
		}
		if tmp < sum {
			left++
			continue
		}
		right--
	}
	return 0, 0
}
