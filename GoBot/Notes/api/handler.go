package api

import (
	"log"
	"time"
)

func UserCreat(tgid int64, name string) string {
	users := User{
		TelegramID: tgid,
		Username:   name,
		CreatedAt:  time.Now(),
	}
	_, err := db.Exec("INSERT INTO Users (telegram_id,username,created_at) VALUES ($1, $2, $3)", users.TelegramID, users.Username, users.CreatedAt)
	if err != nil {
		return "Ваш аккаунт найден"
	}

	return "Успешная регистрация"
}

func NotePlus(tgid int64, title, content string) string {
	var dbUser UserID
	err := db.Get(&dbUser, "SELECT id FROM Users WHERE telegram_id=$1", tgid)
	if err != nil {
		log.Printf("Ошибка при получении user_id: %v\n", err)
		return "Ошибка при получении пользователя"
	}

	newNote := Note{
		UserID:    int(dbUser.Userid),
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	_, err = db.Exec("INSERT INTO Notes (user_id, title, content, created_at, updated_at) VALUES ($1, $2, $3, $4, $5)", newNote.UserID, newNote.Title, newNote.Content, newNote.CreatedAt, newNote.UpdatedAt)
	if err != nil {
		log.Fatalf("ТУТ ОШИБКА %v\n", err.Error())
		return "Ошибка при добавлении заметки"
	}
	return "Заметка добавлена"
}

// func NoteList(tgid int64) string {
// 	var dbUser UserID
// 	err := db.Get(&dbUser, "SELECT id FROM Notes WHERE id=$1", tgid)
// 	if err != nil {
// 		log.Printf("Ошибка при получении user_id: %v\n", err)
// 		return "Ошибка при получении пользователя"
// 	}
// }
