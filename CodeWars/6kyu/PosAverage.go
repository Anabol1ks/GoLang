package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PosAverage(s string) float64 {
	spl := strings.Split(s, ", ")
	var sum, total float64

	for str1 := 0; str1 < len(spl)-1; str1++ {
		for str2 := str1 + 1; str2 < len(spl); str2++ {
			for i := 0; i < len(spl[0]); i++ {
				str_1 := spl[str1]
				str_2 := spl[str2]
				if str_1[i] == str_2[i] {
					sum++
				}
				total++
			}
		}
	}

	r := sum / total * 100
	res, _ := strconv.ParseFloat(fmt.Sprintf("%.10f", r), 10)
	return res
}

func main() {
	fmt.Println(PosAverage("6900690040, 4690606946, 9990494604"))
}
