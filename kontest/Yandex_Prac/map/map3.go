package main

import "fmt"

func main() {
	product := map[string]int{
		"хлеб":     50,
		"молоко":   100,
		"масло":    200,
		"колбаса":  500,
		"соль":     20,
		"огурцы":   200,
		"сыр":      600,
		"ветчина":  700,
		"буженина": 900,
		"помидоры": 250,
		"рыба":     300,
		"хамон":    1500,
	}

	z := []string{"хлеб", "буженина", "сыр", "огурцы"}
	sum_z := 0

	for k, v := range product {
		if v >= 500 {
			fmt.Printf("%v является деликатесом \n", k)
		}
	}
	for _, item := range z {
		sum_z += product[item]
	}
	fmt.Println("сумма продуктов по списку", sum_z)
}
