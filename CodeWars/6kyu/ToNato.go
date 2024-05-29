package main

import (
	"fmt"
	"strings"
)

var NATO = map[string]string{
	"A": "Alfa",
	"B": "Bravo",
	"C": "Charlie",
	"D": "Delta",
	"E": "Echo",
	"F": "Foxtrot",
	"G": "Golf",
	"H": "Hotel",
	"I": "India",
	"J": "Juliett",
	"K": "Kilo",
	"L": "Lima",
	"M": "Mike",
	"N": "November",
	"O": "Oscar",
	"P": "Papa",
	"Q": "Quebec",
	"R": "Romeo",
	"S": "Sierra",
	"T": "Tango",
	"U": "Uniform",
	"V": "Victor",
	"W": "Whiskey",
	"X": "Xray",
	"Y": "Yankee",
	"Z": "Zulu",
}

func ToNato(words string) (res string) {
	words = strings.ToUpper(words)
	words = strings.ReplaceAll(words, " ", "")
	for w, i := range words {
		if string(i) != "!" && string(i) != "?" && string(i) != "." && string(i) != "," {
			res += NATO[string(i)]
			if w != len(words)-1 {
				res += " "
			}
		}
		if string(i) == "!" || string(i) == "?" || string(i) == "." || string(i) == "," {
			res += string(i)
		}
	}
	return
}

func main() {
	fmt.Println(ToNato("If you can read"))
}
