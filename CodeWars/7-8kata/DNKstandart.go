package main

import (
	"fmt"
	"strings"
)

// func DNAStrand(dna string) (dna_plus string) {
// 	for i := range dna {
// 		if dna[i] == 'T' {
// 			dna_plus += "A"
// 		} else if dna[i] == 'A' {
// 			dna_plus += "T"
// 		} else if dna[i] == 'G' {
// 			dna_plus += "C"
// 		} else if dna[i] == 'C' {
// 			dna_plus += "G"
// 		} else {
// 			dna_plus += string(dna[i])
// 		}
// 	}
// 	return
// }

func DNAStrand(dna string) string {
	return strings.NewReplacer("A", "T", "T", "A", "G", "C", "C", "G").Replace(dna)
}

func main() {

	dna := "ATTGC"
	fmt.Println(DNAStrand(dna))
}
