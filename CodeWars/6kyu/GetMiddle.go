package main

import "fmt"

func GetMiddle(s string) string {
	l := len(s)
	if l > 2 {
		if l%2 == 0 {
			return s[l/2-1 : l/2+1]
		}
		return s[l/2 : l/2+1]
	}
	return s
}

func main() {
	fmt.Println(GetMiddle("testing"))
}
