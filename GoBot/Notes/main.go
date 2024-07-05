package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

var db *sqlx.DB

// User структура представляет пользователя.
type User struct {
	ID         int       `json:"id"`
	TelegramID int64     `json:"telegram_id"`
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
}

// Note структура представляет заметку.
type Note struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// Tag структура представляет тег.
type Tag struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// NoteTag структура представляет связь между заметкой и тегом.
type NoteTag struct {
	NoteID int `json:"note_id"`
	TagID  int `json:"tag_id"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	var token = os.Getenv("TELEGRAM_TOKEN")

	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	// userStates := make(map[int64]string)

	for update := range updates {
		if update.Message != nil {
			chatID := update.Message.Chat.ID
			userID := update.Message.From.ID
			userName := update.Message.From.UserName
			userMessage := update.Message.Text
			if userMessage == "рег" {
				id, _ := json.Marshal(map[string]int64{"telegram_id": userID})
				name, _ := json.Marshal(map[string]string{"username": userName})

				msg := tgbotapi.NewMessage(chatID, string(id)+string(name))
				bot.Send(msg)
			}
		}
	}
}
