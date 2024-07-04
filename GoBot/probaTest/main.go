package main

import (
	"encoding/json"
	"fmt"
	"log"
	"proba/weather"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var api = "7374075996:AAExjhESgqe_Nanh34yZA_S2e2AMnSnHnhs"
var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("степень"),
		tgbotapi.NewKeyboardButton("фото"),
		tgbotapi.NewKeyboardButton("погода"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

// var numericKeyboard2 = tgbotapi.NewInlineKeyboardMarkup(
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonURL("1.com", "http://1.com"),
// 		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
// 		tgbotapi.NewInlineKeyboardButtonData("3", "3"),
// 	),
// 	tgbotapi.NewInlineKeyboardRow(
// 		tgbotapi.NewInlineKeyboardButtonData("4", "4"),
// 		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
// 		tgbotapi.NewInlineKeyboardButtonData("6", "6"),
// 	),
// )

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

			if userMessage == "степень" {
				msg := tgbotapi.NewMessage(chatID, "Введите число для возведения в степень")
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
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
			} else if userMessage == "button1" {
				msg := tgbotapi.NewMessage(chatID, "dwd")
				msg.ReplyMarkup = numericKeyboard
				bot.Send(msg)
			} else if userMessage == "close" {
				msg := tgbotapi.NewMessage(chatID, "buttons close")
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				bot.Send(msg)
			} else if userMessage == "фото" {
				file := tgbotapi.FilePath("./gung.jpg")
				msg := tgbotapi.NewPhoto(chatID, file)
				bot.Send(msg)
			} else if userMessage == "погода" {
				userStates[chatID] = "город"
				msg := tgbotapi.NewMessage(chatID, "Введите название города")
				msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
				bot.Send(msg)
			} else if userStates[chatID] == "город" {
				cityJSON, _ := json.Marshal(map[string]string{"name": userMessage})
				city := string(cityJSON)
				fmt.Println(city)
				res := weather.Weather(city)
				msg := tgbotapi.NewMessage(chatID, res)
				bot.Send(msg)
				userStates[chatID] = ""
			}
		}
	}
}
