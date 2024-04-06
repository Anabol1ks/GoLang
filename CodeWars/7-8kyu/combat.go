package main

import "fmt"

func combat(health, damage float64) float64 {
	if health > 0 {
		if health <= damage {
			return 0
		}
		return health - damage
	} else {
		return 0
	}
}

func main() {
	fmt.Println((combat(100, 32.3)))
}
