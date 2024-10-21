package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/mail"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
)

func getTokenFromFile(filePath string) (*oauth2.Token, error) {
	file, err := os.Open((filePath))
	if err != nil {
		return nil, err
	}
	defer file.Close()

	token := &oauth2.Token{}
	err = json.NewDecoder(file).Decode(token)
	return token, err
}

func saveToken(filePath string, token *oauth2.Token) {
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Не удалось сохранить токен в файл: %v", err)
	}
	defer file.Close()
	json.NewEncoder(file).Encode(token)
}

func getClient(config *oauth2.Config) *http.Client {
	tokenFile := "token.json"
	token, err := getTokenFromFile(tokenFile)
	if err != nil {
		token = getTokenFromWeb(config)
		saveToken(tokenFile, token)
	}
	return config.Client(context.Background(), token)
}

func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Перейдите по следующей ссылке для авторизации: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Не удалось считать код авторизации: %v", err)
	}

	token, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Не удалось получить токен доступа: %v", err)
	}
	return token
}

func getSender(headers []*gmail.MessagePartHeader) string {
	for _, header := range headers {
		if header.Name == "From" {
			// Используем пакет net/mail для разбора заголовка From
			addr, err := mail.ParseAddress(header.Value)
			if err != nil {
				log.Printf("Не удалось разобрать адрес отправителя: %v", err)
				return header.Value
			}
			return addr.Address
		}
	}
	return ""
}

func main() {
	creadentialFile := "credentials.json"
	data, err := os.ReadFile(creadentialFile)
	if err != nil {
		log.Fatalf("Не удалось прочитать файл credentials: %v", err)
	}

	config, err := google.ConfigFromJSON(data, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Не удалость создать конфигурацию: %v", err)
	}

	client := getClient(config)
	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Не удалось создать Gmail клиента: %v", err)
	}

	user := "me"
	msgList, err := srv.Users.Messages.List(user).MaxResults(20).Do()
	if err != nil {
		log.Fatalf("Не удалось получить сообщение: %v", err)
	}

	if len(msgList.Messages) == 0 {
		fmt.Println("Сообщений нет")
	} else {
		filteredSenders := []string{
			"online@mirea.ru",
			"oplata-it@mirea.ru",
		}
		fmt.Println("Сообщения от выбранных отправителей: ")
		for _, msg := range msgList.Messages {
			message, err := srv.Users.Messages.Get(user, msg.Id).Do()
			if err != nil {
				log.Printf("Не удалось получить детали сообщения %s: %v", msg.Id, err)
				continue
			}

			sender := getSender(message.Payload.Headers)

			for _, filteredSender := range filteredSenders {
				if sender == filteredSender {
					fmt.Printf("Сообщение от %s: %s\n", sender, message.Snippet)
					break
				}
			}
		}
	}
}
