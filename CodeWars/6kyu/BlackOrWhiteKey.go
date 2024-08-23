package main

import (
	"fmt"
)

func dvb() {
	num := n + 5
	for n < num {
		n++
		if n%2 == 1 {
			kl[n] = "black"
		} else {
			kl[n] = "white"
		}
	}
}

func trb() {
	num := n + 7
	for n < num {
		n++
		if n%2 == 1 {
			kl[n] = "white"
		} else {
			kl[n] = "black"
		}
	}
}
func piano() {
	kl[1], kl[2], kl[3], kl[88] = "white", "black", "white", "white"
	for n < 81 {
		dvb()
		trb()
	}
}

var kl = make(map[int]string)
var n = 3

func BlackOrWhiteKey(keyPressCount int) string {
	piano()
	if keyPressCount > 88 {
		if keyPressCount%88 == 0 {
			return kl[1]
		}
		return kl[keyPressCount%88]
	}
	return kl[keyPressCount]
}

func main() {
	fmt.Println(BlackOrWhiteKey(196))
}
