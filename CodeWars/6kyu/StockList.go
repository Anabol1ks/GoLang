package main

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) (res string) {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}
	for n, c := range listCat {
		res += "(" + c + " : "
		count := 0
		for _, i := range listArt {
			n := strings.Split(i, " ")
			f1 := string(i[0])
			if f1 == c {
				q, _ := strconv.Atoi(n[1])
				count += q
			}
		}
		res += strconv.Itoa(count) + ")"
		if n != len(listCat)-1 {
			res += " - "
		}
	}
	return
}

func main() {
	var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
	var c = []string{"B", "R", "D", "X"}
	fmt.Println(StockList(b, c))

}
