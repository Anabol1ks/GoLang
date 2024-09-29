package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	days := []int{}
	k := map[string]int{
		"MON": 1, "TUE": 2, "WED": 3, "THU": 4, "FRI": 5, "SAT": 6, "SUN": 7,
	}
	kol := make([][]string, 4)
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < 4; i++ {
		scanner.Scan()
		input := scanner.Text()
		elements := strings.Fields(input)

		if len(elements) > 7 {
			elements = elements[:7]
		}
		kol[i] = elements
	}
	for i := 0; i < 4; i++ {
		for j := 0; j < len(kol[i]); j++ {
			el := kol[i][j]
			num, exists := k[el]
			if exists {
				num = num + (7 * i)
				days = append(days, num)
			}
		}
	}
	if len(days) == 28 {
		fmt.Println(0, 0)
		os.Exit(0)
	}
	ras := 0
	d1, d2 := 1, 28
	for i := 0; i < len(days)-1; i++ {
		r := days[i+1] - days[i]
		if ras < r {
			ras = r
			d1 = days[i] + 1
			d2 = days[i+1] - 1
		}
	}
	fmt.Println(d1, d2)
}
