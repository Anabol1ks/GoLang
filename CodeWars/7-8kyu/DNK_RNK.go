package main

import "fmt"

func DNAtoRNA(dna string) string {
	rna := ""
	for i := 0; i < len(dna); i++ {
		if dna[i] == 'T' {
			rna += "U"
		} else {
			rna += string(dna[i])
		}
	}
	return rna
}

func main() {
	s := DNAtoRNA("GCAT")
	fmt.Println(s)
}
