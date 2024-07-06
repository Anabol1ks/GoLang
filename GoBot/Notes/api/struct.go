package api

import (
	"time"

	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func SetDB(database *sqlx.DB) {
	db = database
}

func RunDB() error {
	var err error
	db, err = sqlx.Open("postgres", "host=localhost port=2222 user=postgres password=qwerty dbname=notesbot sslmode=disable")
	if err != nil {
		return err
	}

	// Проверяем соединение с базой данных
	err = db.Ping()
	if err != nil {
		db.Close()
		return err
	}

	SetDB(db)
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

// User структура представляет пользователя.
type User struct {
	ID         int       `json:"id"`
	TelegramID int64     `json:"telegram_id"`
	Username   string    `json:"username"`
	CreatedAt  time.Time `json:"created_at"`
}

type UserID struct {
	Userid int64 `db:"id"`
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
