package main

import "fmt"

func Revrot(s string, n int) (res string) {
	if len(s) < n || n == 0 {
		return ""
	}
	sum := 0
	str := ""
	for i, ch := range s {
		sum += int(ch) - 48
		str += string(ch)
		if (i+1)%n == 0 {
			if sum%2 == 0 {
				for j := len(str) - 1; j >= 0; j-- {
					res += string(str[j])
				}
			} else {
				res += str[1:len(str)] + str[:1]
			}
			sum = 0
			str = ""
		}
	}
	return
}

func main() {
	fmt.Println(Revrot("2246875", 5))
}
