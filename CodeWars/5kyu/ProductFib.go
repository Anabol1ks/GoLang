package main

import "fmt"

func ProductFib(prod uint64) (res [3]uint64) {
	m := []int{}
	m = append(m, 0, 1)
	for i := 0; i < int(prod); i++ {
		m = append(m, (m[i] + m[i+1]))
		var s uint64 = uint64(m[i] * m[i+1])
		if s >= prod {
			res[0], res[1] = uint64(m[i]), uint64(m[i+1])
			if s == prod {
				res[2] = 1
			}
			break
		}
	}

	return
}

func main() {
	fmt.Println(ProductFib(5895))

}
