package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go printNumbers() // Запуск функции в goroutine
	fmt.Println("Goroutine запущена!")

	// Подождём немного, чтобы увидеть результат работы goroutine
	time.Sleep(6 * time.Second)
	fmt.Println("Завершение работы программы.")
}
