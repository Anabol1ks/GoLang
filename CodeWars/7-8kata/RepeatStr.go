package main

import "fmt"

func RepeatStr(repetitions int, value string) (s string) {
	for i := 1; i <= repetitions; i++ {
		s += fmt.Sprint(value)
	}
	return
}

func main() {
	fmt.Println(RepeatStr(4, "d"))
}
