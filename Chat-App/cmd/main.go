package main

import (
	"chat_app/internal/auth"
	"chat_app/internal/chat"
	"chat_app/internal/storage"
	"chat_app/internal/ws"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env")
	}

	storage.ConnectDatabase()

	err = storage.DB.AutoMigrate(&auth.User{}, &chat.Message{}, &chat.Room{})
	if err != nil {
		log.Fatal("Ошибка миграции", err)
	}

	r := gin.Default()

	r.POST("auth/register", auth.RegisterHandler)
	r.POST("auth/login", auth.LoginHandler)

	go ws.HandleMessage()
	r.GET("/ws", ws.HandleConnections)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера")
	}
}