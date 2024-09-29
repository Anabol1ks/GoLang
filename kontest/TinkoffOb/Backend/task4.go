package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var platforms [][]int
	for i := 0; i < m; i++ {
		var s, t int
		fmt.Scan(&s, &t)
		platforms = append(platforms, []int{s, t})
	}

	prevStart, prevEnd := platforms[0][0], platforms[0][1]
	actions := 0

	for i := 1; i < m; i++ {
		currStart, currEnd := platforms[i][0], platforms[i][1]

		if currStart <= prevEnd && currEnd >= prevStart {
		} else {
			if currStart > prevEnd {
				actions += currStart - prevEnd
				currStart = prevEnd
				currEnd = currStart + (platforms[i][1] - platforms[i][0])
			} else if currEnd < prevStart {
				actions += prevStart - currEnd
				currEnd = prevStart
				currStart = currEnd - (platforms[i][1] - platforms[i][0])
			}
		}

		prevStart, prevEnd = currStart, currEnd
	}

	fmt.Println(actions)
}
