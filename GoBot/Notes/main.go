package main

import (
	"log"
	"notes/api"
	"os"
	"strconv"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var Keyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Список"),
		tgbotapi.NewKeyboardButton("Новая запись"),
	),
)

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
	err = api.RunDB()
	if err != nil {
		log.Fatalf("Ошибка при запуске базы данных: %v", err)
	}
	defer api.CloseDB()
	userStates := make(map[int64]string)
	var title, content string

	for update := range updates {
		if update.Message == nil {
			continue
		}
		chatID := update.Message.Chat.ID
		state, _ := userStates[chatID]
		userID := update.Message.From.ID
		userName := update.Message.From.UserName
		userMessage := update.Message.Text
		switch userMessage {
		case "/start":
			create := api.UserCreat(userID, userName)
			msg := tgbotapi.NewMessage(chatID, create)
			bot.Send(msg)
			time.Sleep(50 * time.Millisecond)
			msg = tgbotapi.NewMessage(chatID, "Выберите дальнейшее действие")
			msg.ReplyMarkup = Keyboard
			bot.Send(msg)
		case "id":
			msg := tgbotapi.NewMessage(chatID, strconv.Itoa(int(chatID)))
			bot.Send(msg)
		case "Новая запись":
			msg := tgbotapi.NewMessage(chatID, "Введите название заголовка")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
			userStates[chatID] = "титл"
		case "Список":
			msg := tgbotapi.NewMessage(chatID, "Список ваших заметок")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
			api.StrNotes(userID, bot, chatID)
		case "Выбрать заметку":
			msg := tgbotapi.NewMessage(chatID, "Введите название заметки")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			userStates[chatID] = "вывод"
			bot.Send(msg)
		case "Изменить заметку":
			msg := tgbotapi.NewMessage(chatID, "Введите название заметки")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			userStates[chatID] = "редакт"
			bot.Send(msg)
		case "Удалить заметку":
			userStates[chatID] = "delete"
			msg := tgbotapi.NewMessage(chatID, "Введите название заметки")
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
			bot.Send(msg)
		}
		if state == "титл" {
			title = userMessage
			userStates[chatID] = "заметка"
			msg := tgbotapi.NewMessage(chatID, "Напишите саму заметку")
			bot.Send(msg)
		}
		if state == "заметка" {
			content = userMessage
			msg := tgbotapi.NewMessage(chatID, "Получил")
			bot.Send(msg)
			time.Sleep(1 * time.Second)
			notes := api.NotePlus(userID, title, content)
			msg = tgbotapi.NewMessage(chatID, notes)
			bot.Send(msg)
			userStates[chatID] = ""
			title, content = "", ""
		}
		if state == "delete" {
			title = userMessage
			del := api.DelNote(chatID, title)
			msg := tgbotapi.NewMessage(chatID, del)
			msg.ReplyMarkup = Keyboard
			bot.Send(msg)
			title = ""
			userStates[chatID] = ""
		}
		if state == "редакт" {
			title = userMessage
			msg := tgbotapi.NewMessage(chatID, "Введите новое содеражание заметки")
			bot.Send(msg)
			userStates[chatID] = "редакт+"
		}
		if state == "редакт+" {
			content = userMessage
			userStates[chatID] = ""
			res := api.UpdateNote(chatID, title, content)
			msg := tgbotapi.NewMessage(chatID, res)
			msg.ReplyMarkup = Keyboard
			bot.Send(msg)
			content = ""
			title = ""
		}
		if state == "вывод" {
			title = userMessage
			contentPr := api.PrintNote(chatID, title)
			msg := tgbotapi.NewMessage(chatID, contentPr)
			msg.ReplyMarkup = Keyboard
			bot.Send(msg)
			title = ""
			userStates[chatID] = ""
		}
	}
}
