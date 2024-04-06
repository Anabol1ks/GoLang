package main

import (
	"fmt"
)

func RemoveChar(word string) string {
	return word[1 : len(word)-1]
}

func main() {
	str := "eloquent"
	fmt.Println(RemoveChar(str))

}
