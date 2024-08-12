package main

import (
	"fmt"
)

func sendData(ch chan int) {
	ch <- 10 // Отправляем значение в канал
}

func main() {
	ch := make(chan int) // Создаём канал, который передает целые числа

	go sendData(ch) // Запускаем goroutine, которая отправит данные в канал

	receivedValue := <-ch // Ожидаем получения данных из канала
	fmt.Println("Получено значение:", receivedValue)
}
