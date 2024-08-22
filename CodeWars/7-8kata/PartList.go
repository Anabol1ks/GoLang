package main

import "fmt"

func PartList(arr []string) (res string) {
	for i := 0; i < len(arr)-1; i++ {
		res += "("
		for n, str := range arr {
			if n == i {
				res += str + "," + " "
			} else {
				res += str + " "
			}
		}
		res = res[:len(res)-1] + ")"
	}

	return
}

func main() {
	fmt.Println(PartList([]string{"I", "wish", "I", "hadn't", "come"}))
}
