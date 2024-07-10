package api

import (
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/lib/pq"
)

var Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Список"),
		tgbotapi.NewKeyboardButton("Новая запись"),
	),
)
var Sps = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Выбрать заметку"),
		tgbotapi.NewKeyboardButton("Изменить заметку"),
		tgbotapi.NewKeyboardButton("Удалить заметку"),
	),
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
		if pqErr, ok := err.(*pq.Error); ok {
			if pqErr.Code.Name() == "unique_violation" {
				log.Printf("У вас уже существует заметка с таким названием: %v\n", err)
				return "У вас уже существует заметка с таким названием"
			}
		}
		log.Printf("Ошибка при добавлении заметки: %v\n", err)
		return "Ошибка при добавлении заметки"
	}
	return "Заметка добавлена"
}

func NoteList(tgid int64) ([]string, error) {
	var dbUser UserID
	err := db.Get(&dbUser, "SELECT id FROM Users WHERE telegram_id=$1", tgid)
	if err != nil {
		log.Printf("Ошибка при получении user_id: %v\n", err)
		return nil, fmt.Errorf("Ошибка при получении пользователя")
	}
	var titles []string

	rows, err := db.Query("SELECT title FROM Notes WHERE user_id=$1", dbUser.Userid)
	if err != nil {
		log.Printf("Ошибка при получении заголовков: %v\n", err)
		return nil, fmt.Errorf("Ошибка при получении заголовков")
	}
	defer rows.Close()

	for rows.Next() {
		var title string
		if err := rows.Scan(&title); err != nil {
			log.Printf("Ошибка при чтении заголовка: %v\n", err)
			return nil, fmt.Errorf("Ошибка при чтении заголовка")
		}
		titles = append(titles, title)
	}
	return titles, nil
}

func StrNotes(tgid int64, bot *tgbotapi.BotAPI, chatID int64) {
	titles, err := NoteList(tgid)
	if err != nil {
		log.Printf("Ошибка при получении заголовков: %v\n", err)
		msg := tgbotapi.NewMessage(chatID, "Ошибка при получении заголовков")
		bot.Send(msg)
		return
	}
	if len(titles) == 0 {
		msg := tgbotapi.NewMessage(chatID, "Нет заметок для этого пользователя")
		msg.ReplyMarkup = Keyboard
		bot.Send(msg)
		return
	}
	for i := range titles {
		titles[i] = " · " + titles[i]
	}
	titlesStr := strings.Join(titles, "\n")
	msg := tgbotapi.NewMessage(chatID, fmt.Sprintf("Заголовки заметок:\n%s", titlesStr))
	msg.ReplyMarkup = Sps
	bot.Send(msg)
}

func PrintNote(tgid int64, title string) (string, error) {
	var dbUser UserID
	err := db.Get(&dbUser, "SELECT id FROM Users WHERE telegram_id=$1", tgid)
	if err != nil {
		log.Printf("Ошибка при получении user_id: %v\n", err)
		return "", fmt.Errorf("не удалось найти пользователя с указанным telegram_id")
	}
	var content Content
	err = db.QueryRowx("SELECT content FROM notes WHERE user_id=$1 AND title=$2", dbUser.Userid, title).StructScan(&content)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("заметка с указанным заголовком не найдена")
		}
		log.Printf("Ошибка при получении содержания заметки: %v\n", err)
		return "", fmt.Errorf("ошибка при получении содержания заметки")
	}

	return content.Content, nil
}

func UpdateNote(tgid int64, title, text string) (string, error) {
	var dbUser UserID
	err := db.Get(&dbUser, "SELECT id FROM Users WHERE telegram_id=$1", tgid)
	if err != nil {
		log.Printf("Ошибка при получении user_id: %v\n", err)
		return "", fmt.Errorf("не удалось найти пользователя с указанным telegram_id")
	}
	timeUp := time.Now()
	oldcontent, err := PrintNote(tgid, title)
	if err != nil {
		log.Printf("Ошибка при получении содержания заметки: %v\n", err)
		return "", fmt.Errorf("Нет заметки с таким названием")
	}

	_, err = db.Exec("UPDATE notes SET content=$1, updated_at=$2 WHERE user_id=$3 AND title=$4", text, timeUp, dbUser.Userid, title)
	if err != nil {
		log.Printf("Ошибка при обновлении заметки: %v\n", err)
		return "", fmt.Errorf("не удалось обновить заметку. Возможно, заметка с таким заголовком не существует")
	}
	res := fmt.Sprintf(
		"Информация в заметке изменена с:\n · %s\nНа:\n · %s",
		oldcontent, text,
	)
	return res, nil
}

func DelNote(tgid int64, title string) string {
	var dbUser UserID
	err := db.Get(&dbUser, "SELECT id FROM Users WHERE telegram_id=$1", tgid)
	if err != nil {
		log.Printf("Ошибка при получении user_id: %v\n", err)
		return "Ошибка при получении пользователя"
	}
	result, err := db.Exec("DELETE FROM notes WHERE user_id=$1 AND title=$2", dbUser.Userid, title)
	if err != nil {
		log.Printf("Ошибка при удалении заметки: %v\n", err)
		return "Ошибка при удалении заметки"
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Ошибка при получении количества затронутых строк: %v\n", err)
		return "Ошибка при удалении заметки"
	}

	if rowsAffected == 0 {
		return "Заметка с таким заголовком не найдена"
	}

	return "Заметка " + title + " успешно удалена"
}
