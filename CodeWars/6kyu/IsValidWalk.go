package main

import "fmt"

func IsValidWalk(walk []rune) bool {
	if len(walk) == 10 {
		var x, y int
		for _, i := range walk {
			switch i {
			case 'n':
				x++
			case 's':
				x--
			case 'e':
				y++
			case 'w':
				y--
			}
		}
		if x == 0 && y == 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(IsValidWalk([]rune{'n', 's', 'n', 's', 'n', 's', 'n', 's', 'n', 's'}))
}
