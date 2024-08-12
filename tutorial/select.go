package main

import (
	"fmt"
	"time"
)

func sendMessage(ch chan string, message string, delay time.Duration) {
	time.Sleep(delay)
	ch <- message
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go sendMessage(ch1, "Сообщение из ch1", 2*time.Second)
	go sendMessage(ch2, "Сообщение из ch2", 1*time.Second)

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println(msg1)
		case msg2 := <-ch2:
			fmt.Println(msg2)
		}
	}
}
