package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}
	port := os.Getenv("PORT")

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Привет, всё ок"})
	})

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Ошибка запуска")
	}
}
