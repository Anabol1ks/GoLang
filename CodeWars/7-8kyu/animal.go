package main

import "fmt"

func Feast(beast string, dish string) bool {
	BeastOne := beast[0]
	BeastLast := beast[len(beast)-1]
	DishOne := dish[0]
	DishLast := dish[len(dish)-1]
	var b bool
	if BeastOne == DishOne && BeastLast == DishLast {
		b = true
	} else {
		b = false
	}
	return b
}

// func main() {
// 	name := "Ananas"
// 	s := name[len(name)-1]
// 	fmt.Println(string(s))
// }

func main() {
	f := Feast("Chicate", "cinate")
	fmt.Println(f)
}
