package main

import (
	"fmt"
	"sort"
	"strings"
)

type NameWeight struct {
	name   string
	weight int
}

func NthRank(st string, we []int, n int) string {
	if st == "" {
		return "No participants"
	}

	names := strings.Split(st, ",")
	if n > len(names) {
		return "Not enough participants"
	}

	al := make(map[rune]int)
	ves := 1
	for i := 'a'; i <= 'z'; i++ {
		al[i] = ves
		ves++
	}

	var nameWeights []NameWeight

	for i, str := range names {
		sum := 0
		for _, ch := range strings.ToLower(str) {
			sum += al[ch]
		}
		sum += len(str)
		sum *= we[i]

		nameWeights = append(nameWeights, NameWeight{name: str, weight: sum})
	}

	sort.Slice(nameWeights, func(i, j int) bool {
		if nameWeights[i].weight == nameWeights[j].weight {
			return nameWeights[i].name < nameWeights[j].name
		}
		return nameWeights[i].weight > nameWeights[j].weight
	})

	return nameWeights[n-1].name
}

func main() {
	fmt.Println(NthRank("COLIN,AMANDBA,AMANDAB,CAROL,PauL,JOSEPH", []int{1, 4, 4, 5, 2, 1}, 4))
}
