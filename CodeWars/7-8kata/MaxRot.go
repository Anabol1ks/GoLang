package main

import (
	"fmt"
	"strconv"
)

func MaxRot(n int64) int64 {
	str := strconv.FormatInt(n, 10)
	max := n

	for i := 0; i < len(str)-1; i++ {
		str = str[:i] + str[i+1:] + string(str[i])
		num, _ := strconv.ParseInt(str, 10, 64)
		if max < num {
			max = num
		}
	}

	return max
}

func main() {
	fmt.Println(MaxRot(56789))
}
