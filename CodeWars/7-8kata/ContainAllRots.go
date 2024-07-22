package main

import "fmt"

func getAllRotations(strng string) []string {
	n := len(strng)
	rotations := make([]string, n)
	for i := 0; i < n; i++ {
		rotations[i] = strng[i:] + strng[:i]
	}
	return rotations
}

func ContainAllRots(strng string, arr []string) bool {
	if strng == "" {
		return true
	}
	rotations := getAllRotations(strng)
	for _, rotation := range rotations {
		if !contains(arr, rotation) {
			return false
		}
	}
	return true
}

func contains(arr []string, s string) bool {
	for _, a := range arr {
		if a == s {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(ContainAllRots("XjYABhR", []string{"TzYxlgfnhf", "yqVAuoLjMLy", "BhRXjYA", "YABhRXj", "hRXjYAB", "jYABhRX", "XjYABhR", "ABhRXjY"}))
}
