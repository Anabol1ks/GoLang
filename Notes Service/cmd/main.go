package main

import (
	"log"
	"notes-service/internal/auth"
	"notes-service/internal/notes"
	"notes-service/internal/storage"
	"os"

	_ "notes-service/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Ошибка загрузки .env файла")
	}

	storage.ConnectDatabase()

	err = storage.DB.AutoMigrate(&notes.Note{}, &auth.User{})
	if err != nil {
		log.Fatal("Ошибка миграции!", err.Error())
	}
	port := os.Getenv("PORT")

	c := cron.New()
	c.AddFunc("@daily", func() {
		notes.CleanupDeletedNotes()
	})
	c.Start()

	defer c.Stop()

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.POST("/auth/register", auth.RegisterHandler)
	r.POST("/auth/login", auth.LoginHandler)

	// Группа маршрутов с авторизацией
	authorized := r.Group("/")
	authorized.Use(auth.AuthMiddleware())
	authorized.GET("/notes", notes.GetNotesHandler)
	authorized.POST("/notes", notes.CreateNoteHandler)
	authorized.GET("/notes/:id", notes.GetNoteHandler)
	authorized.PUT("/notes/:id", notes.UpdateNoteHandler)
	authorized.DELETE("/notes/:id", notes.DeleteNoteHandler)

	err = r.Run(":" + port)
	if err != nil {
		log.Fatal("Ошибка запуска")
	}
}
