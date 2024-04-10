package main

import "fmt"

func GetVolumeOfCuboid(length, width, height float64) float64 {
	return length * width * height
}

func main() {
	fmt.Println(GetVolumeOfCuboid(6.3, 2.0, 5.0))
}
