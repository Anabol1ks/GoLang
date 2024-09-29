package main

import "fmt"

func main() {
	m := make(map[string]string) // создаём map — о применении функции make для создания переменных типа map будет рассказано ниже
	m["foo"] = "bar"             // заполняем парами «ключ-значение»
	m["ping"] = "pong"
	fmt.Println(m) // печатаем
}
