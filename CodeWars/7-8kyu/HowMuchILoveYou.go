package main

import "fmt"

func HowMuchILoveYou(i int) string {
	if i >= 0 {
		i = i % 6
		if i == 0 {
			i = 6
		}
		switch i {
		case 1:
			return "I love you"
		case 2:
			return "a little"
		case 3:
			return "a lot"
		case 4:
			return "passionately"
		case 5:
			return "madly"
		case 6:
			return "not at all"
		}
	}
	return ""
}

func main() {
	fmt.Println(HowMuchILoveYou(6))
}
