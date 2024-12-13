package main

import (
	_ "chat_app/docs"
	"chat_app/internal/auth"
	"chat_app/internal/chat"
	"chat_app/internal/storage"
	"chat_app/internal/ws"
	"log"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("auth/register", auth.RegisterHandler)
	r.POST("auth/login", auth.LoginHandler)

	authorized := r.Group("/")
	authorized.Use(auth.AuthMiddleware())
	authorized.POST("/rooms", chat.CreateRoomHandler)
	authorized.GET("/rooms", chat.GetRoomsHandler)
	authorized.GET("/rooms/:room_id/messages", chat.GetRoomMessage)
	authorized.GET("/ws/:room_id", ws.HandleConnections)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Ошибка запуска сервера")
	}
}
