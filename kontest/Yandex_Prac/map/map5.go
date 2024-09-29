package main

import "fmt"

func RemoveDuplicates(input []string) []string {
	enc := map[string]bool{}
	result := []string{}

	for _, v := range input {
		if enc[v] == false {
			enc[v] = true
			result = append(result, v)
		}
	}
	return result
}

func main() {
	input := []string{
		"cat",
		"dog",
		"bird",
		"dog",
		"parrot",
		"cat",
	}
	result := RemoveDuplicates(input)
	fmt.Println(result)
}
