package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type Person struct {
	Name        string    `json:"Имя"`
	Email       string    `json:"Почта"`
	DateOfBirth time.Time `json:"-"`
}

func main() {
	p := Person{
		Name:        "Alex",
		Email:       "a@yandex.ru",
		DateOfBirth: time.Now(),
	}
	jsMan, err := json.Marshal(p)
	if err != nil {
		log.Fatalln("unable marshal to json")
	}
	fmt.Printf("Man %v", string(jsMan))
}
