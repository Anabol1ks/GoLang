package main

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"

	"github.com/gin-gonic/gin"
)

type Email struct {
	Sender  string `json:"sender"`
	Snippet string `json:"snippet"`
}

type PageVariables struct {
	Emails []Email
}

func getTokenFromFile(filePath string) (*oauth2.Token, error) {
	file, err := os.Open(filePath)
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
	credentialsFile := "credentials.json"
	data, err := ioutil.ReadFile(credentialsFile)
	if err != nil {
		log.Fatalf("Не удалось прочитать файл credentials: %v", err)
	}

	config, err := google.ConfigFromJSON(data, gmail.GmailReadonlyScope)
	if err != nil {
		log.Fatalf("Не удалось создать конфигурацию из файла credentials: %v", err)
	}

	client := getClient(config)

	srv, err := gmail.New(client)
	if err != nil {
		log.Fatalf("Не удалось создать Gmail клиент: %v", err)
	}

	user := "me"
	msgList, err := srv.Users.Messages.List(user).MaxResults(25).Do()
	if err != nil {
		log.Fatalf("Не удалось получить сообщения: %v", err)
	}

	var emails []Email

	if len(msgList.Messages) == 0 {
		fmt.Println("Сообщений нет.")
	} else {
		filteredSenders := []string{
			"online@mirea.ru",
		}

		for _, msg := range msgList.Messages {
			message, err := srv.Users.Messages.Get(user, msg.Id).Do()
			if err != nil {
				log.Printf("Не удалось получить детали сообщения %s: %v", msg.Id, err)
				continue
			}

			sender := getSender(message.Payload.Headers)

			for _, filteredSender := range filteredSenders {
				if sender == filteredSender {
					emails = append(emails, Email{
						Sender:  sender,
						Snippet: message.Snippet,
					})
					break
				}
			}
		}
	}

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		pageVariables := PageVariables{
			Emails: emails,
		}
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			c.String(500, err.Error())
			return
		}
		err = tmpl.Execute(c.Writer, pageVariables)
		if err != nil {
			c.String(500, err.Error())
		}
	})

	log.Println("Сервер запущен на :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
