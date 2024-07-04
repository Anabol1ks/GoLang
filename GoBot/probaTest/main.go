package main

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var api = "7374075996:AAExjhESgqe_Nanh34yZA_S2e2AMnSnHnhs"

func d2(num int) string {
	return strconv.Itoa(num * num)
}

func main() {
	bot, err := tgbotapi.NewBotAPI(api)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	userStates := make(map[int64]string)

	for update := range updates {
		if update.Message != nil {
			log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

			chatID := update.Message.Chat.ID
			userMessage := update.Message.Text

			state, _ := userStates[chatID]

			if userMessage == "/start" {
				msg := tgbotapi.NewMessage(chatID, "Введите число для возведения в степень")
				bot.Send(msg)
				userStates[chatID] = "ожидается число"
			} else if state == "ожидается число" {
				num, err := strconv.Atoi(userMessage)
				if err != nil {
					msg := tgbotapi.NewMessage(chatID, "Введите корректно число")
					bot.Send(msg)
				} else {
					msg := tgbotapi.NewMessage(chatID, " Ваше число в квадрате: "+d2(num))
					bot.Send(msg)
					userStates[chatID] = ""
				}
			}

		}
	}
}
