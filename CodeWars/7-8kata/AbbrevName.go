package main

import (
	"fmt"
	"strings"
)

func AbbrevName(name string) string {
	s := strings.Split(name, " ")
	return strings.ToUpper(string(s[0][0])) + "." + strings.ToUpper(string(s[1][0]))
}

func main() {
	fmt.Println(AbbrevName("patrick feeney"))
}
