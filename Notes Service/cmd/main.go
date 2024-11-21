package main

import (
	"log"
	"notes-service/internal/notes"
	"notes-service/internal/storage"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	storage.ConnectDatabase()

	err = storage.DB.AutoMigrate(&notes.Note{})
	if err != nil {
		log.Fatal("Ошибка миграции!", err.Error())
	}
	port := os.Getenv("PORT")

	r := gin.Default()

	r.POST("/notes", notes.CreateNoteHandler)
	r.GET("/notes", notes.GetNotesHandler)
	r.GET("/notes/:id", notes.GetNoteHandler)
	r.PUT("/notes/:id", notes.UpdateNoteHandler)
	r.DELETE("/notes/:id", notes.DeleteNoteHandler)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Ошибка запуска")
	}
}
