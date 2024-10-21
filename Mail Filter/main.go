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
	"sync"
	"time"

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

type CachedEmails struct {
	Emails     []Email `json:"emails"`
	LastUpdate int64   `json:"last_update"` // Время последнего обновления в Unix формате
}

var cacheMutex sync.Mutex

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

type SenderList struct {
	Senders []string `json:"senders"`
}

func loadSenders(filename string) (SenderList, error) {
	var senderList SenderList
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return SenderList{Senders: []string{}}, nil
		}
		return senderList, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&senderList)
	return senderList, err
}

func saveSenders(filename string, senderList SenderList) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	err = json.NewEncoder(file).Encode(senderList)
	return err
}

// Кеширование писем
func loadCachedEmails() (CachedEmails, error) {
	var cache CachedEmails
	file, err := os.Open("emails_cache.json")
	if err != nil {
		if os.IsNotExist(err) {
			return CachedEmails{}, nil
		}
		return cache, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&cache)
	return cache, err
}

func saveCachedEmails(emails []Email) error {
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	cache := CachedEmails{
		Emails:     emails,
		LastUpdate: time.Now().Unix(),
	}
	file, err := os.Create("emails_cache.json")
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewEncoder(file).Encode(cache)
}

// Асинхронная загрузка писем
func fetchEmailsAsync(srv *gmail.Service, user string, senderList SenderList, resultChan chan<- []Email) {
	msgList, err := srv.Users.Messages.List(user).MaxResults(25).Do()
	if err != nil {
		log.Printf("Ошибка при получении сообщений: %v", err)
		resultChan <- nil
		return
	}

	var emails []Email
	for _, msg := range msgList.Messages {
		message, err := srv.Users.Messages.Get(user, msg.Id).Do()
		if err != nil {
			log.Printf("Ошибка при получении деталей сообщения: %v", err)
			continue
		}

		sender := getSender(message.Payload.Headers)
		for _, filteredSender := range senderList.Senders {
			if sender == filteredSender {
				emails = append(emails, Email{
					Sender:  sender,
					Snippet: message.Snippet,
				})
				break
			}
		}
	}
	resultChan <- emails
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

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		cache, err := loadCachedEmails()
		if err != nil {
			log.Printf("Ошибка при чтении кеша: %v", err)
		}

		// Если кеш обновлялся более 5 минут назад, загружаем новые данные асинхронно
		if time.Now().Unix()-cache.LastUpdate > 5*60 {
			senderList, err := loadSenders("senders.json")
			if err != nil {
				c.String(500, "Ошибка загрузки списка отправителей")
				return
			}

			resultChan := make(chan []Email)
			go fetchEmailsAsync(srv, "me", senderList, resultChan)

			go func() {
				emails := <-resultChan
				if emails != nil {
					saveCachedEmails(emails)
				}
			}()
		}

		pageVariables := PageVariables{
			Emails: cache.Emails,
		}
		tmpl, err := template.ParseFiles("index.html")
		if err != nil {
			c.String(500, err.Error())
			return
		}
		tmpl.Execute(c.Writer, pageVariables)
	})

	r.POST("/senders", func(c *gin.Context) {
		var newSender struct {
			Email string `json:"email"`
		}

		if err := c.ShouldBindJSON(&newSender); err != nil {
			c.JSON(400, gin.H{"error": "Некорректные данные"})
			return
		}

		senderList, err := loadSenders("senders.json")
		if err != nil {
			c.JSON(500, gin.H{"error": "Ошибка чтения файла"})
			return
		}

		// Добавляем нового отправителя в список
		senderList.Senders = append(senderList.Senders, newSender.Email)

		// Сохраняем обновлённый список
		if err := saveSenders("senders.json", senderList); err != nil {
			c.JSON(500, gin.H{"error": "Ошибка сохранения файла"})
			return
		}

		c.JSON(200, gin.H{"status": "Отправитель добавлен"})
	})

	// Удаление отправителя
	r.DELETE("/senders", func(c *gin.Context) {
		var delSender struct {
			Email string `json:"email"`
		}

		if err := c.ShouldBindJSON(&delSender); err != nil {
			c.JSON(400, gin.H{"error": "Некорректные данные"})
			return
		}

		senderList, err := loadSenders("senders.json")
		if err != nil {
			c.JSON(500, gin.H{"error": "Ошибка чтения файла"})
			return
		}

		// Удаляем отправителя из списка
		for i, sender := range senderList.Senders {
			if sender == delSender.Email {
				senderList.Senders = append(senderList.Senders[:i], senderList.Senders[i+1:]...)
				break
			}
		}

		// Сохраняем обновлённый список
		if err := saveSenders("senders.json", senderList); err != nil {
			c.JSON(500, gin.H{"error": "Ошибка сохранения файла"})
			return
		}

		c.JSON(200, gin.H{"status": "Отправитель удалён"})
	})

	log.Println("Сервер запущен на :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
