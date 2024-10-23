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
	SenderEmail string `json:"sender_email"`
	SenderName  string `json:"sender_name"`
	Snippet     string `json:"snippet"`
}

type PageVariables struct {
	Emails  []Email
	Senders []string
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

func getSender(headers []*gmail.MessagePartHeader) (string, string) {
	for _, header := range headers {
		if header.Name == "From" {
			addr, err := mail.ParseAddress(header.Value)
			if err != nil {
				log.Printf("Не удалось разобрать адрес отправителя: %v", err)
				return "", ""
			}
			return addr.Address, addr.Name
		}
	}
	return "", ""
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
func loadCachedEmails(senderList *SenderList) (CachedEmails, error) {
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
	if err != nil {
		return cache, err
	}

	// Фильтруем письма, исключая тех отправителей, которые есть в blackSenders.json
	var filteredEmails []Email
	for _, email := range cache.Emails {
		shouldSkip := false
		for _, sender := range senderList.Senders {
			if email.SenderEmail == sender {
				shouldSkip = true
				break
			}
		}
		if !shouldSkip {
			filteredEmails = append(filteredEmails, email)
		}
	}

	cache.Emails = filteredEmails
	return cache, nil
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
	msgList, err := srv.Users.Messages.List(user).MaxResults(100).Do()
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

		senderEmail, senderName := getSender(message.Payload.Headers)

		// Проверяем, есть ли отправитель в черном списке
		shouldSkip := false
		for _, filteredSender := range senderList.Senders {
			if senderEmail == filteredSender {
				shouldSkip = true
				break
			}
		}

		if !shouldSkip {
			emails = append(emails, Email{
				SenderEmail: senderEmail,
				SenderName:  senderName,
				Snippet:     message.Snippet,
			})
		}
	}

	resultChan <- emails
}

// Автоматическое обновление писем каждые n секунд
func autoUpdateEmails(srv *gmail.Service, senderList *SenderList, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C

		log.Println("Автоматическое обновление писем...")

		resultChan := make(chan []Email)
		go fetchEmailsAsync(srv, "me", *senderList, resultChan)

		go func() {
			emails := <-resultChan
			if emails != nil {
				saveCachedEmails(emails)
				log.Println("Письма обновлены")
			} else {
				log.Println("Ошибка при обновлении писем")
			}
		}()
	}
}

func updateSenderListPeriodically(filePath string, interval time.Duration, senderList *SenderList) {
	for {
		updatedList, err := loadSenders(filePath)
		if err != nil {
			log.Printf("Ошибка загрузки списка отправителей: %v", err)
		} else {
			// Обновляем данные списка отправителей
			*senderList = updatedList
			log.Println("Список отправителей обновлен:", senderList.Senders)
		}

		// Ждем перед следующей проверкой
		time.Sleep(interval)
	}
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

	// Загружаем список отправителей
	senderList, err := loadSenders("blackSenders.json")
	if err != nil {
		log.Fatalf("Ошибка загрузки списка отправителей: %v", err)
	}

	// Запускаем горутину для периодического обновления списка отправителей (через указатель)
	go updateSenderListPeriodically("blackSenders.json", 2*time.Second, &senderList)

	// Запускаем автоматическое обновление писем
	go autoUpdateEmails(srv, &senderList, 5*time.Second)

	r := gin.Default()
	r.StaticFS("/static", http.Dir("static"))
	r.LoadHTMLGlob("./templates/*.html")
	r.GET("/emails", func(c *gin.Context) {
		cache, err := loadCachedEmails(&senderList)
		if err != nil {
			log.Printf("Ошибка при чтении кеша: %v", err)
		}

		pageVariables := PageVariables{
			Emails:  cache.Emails,
			Senders: senderList.Senders,
		}

		tmpl, err := template.ParseFiles("templates/emails.html")
		if err != nil {
			log.Fatalf("Ошибка при парсинге шаблона: %v", err)
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

		// Добавляем нового отправителя
		senderList.Senders = append(senderList.Senders, newSender.Email)
		if err := saveSenders("blackSenders.json", senderList); err != nil {
			c.JSON(500, gin.H{"error": "Ошибка сохранения файла"})
			return
		}

		// Перезагружаем список отправителей после обновления
		updatedList, err := loadSenders("blackSenders.json")
		if err != nil {
			c.JSON(500, gin.H{"error": "Ошибка загрузки обновлённого списка отправителей"})
			return
		}
		senderList = updatedList

		c.JSON(200, gin.H{"status": "Отправитель добавлен"})
	})

	r.DELETE("/senders", func(c *gin.Context) {
		var delSender struct {
			Email string `json:"email"`
		}

		if err := c.ShouldBindJSON(&delSender); err != nil {
			c.JSON(400, gin.H{"error": "Некорректные данные"})
			return
		}

		// Удаляем отправителя
		for i, sender := range senderList.Senders {
			if sender == delSender.Email {
				senderList.Senders = append(senderList.Senders[:i], senderList.Senders[i+1:]...)
				break
			}
		}

		// Сохраняем обновлённый список
		if err := saveSenders("blackSenders.json", senderList); err != nil {
			c.JSON(500, gin.H{"error": "Ошибка сохранения файла"})
			return
		}

		// Перезагружаем список отправителей
		updatedList, err := loadSenders("blackSenders.json")
		if err != nil {
			c.JSON(500, gin.H{"error": "Ошибка загрузки обновлённого списка отправителей"})
			return
		}
		senderList = updatedList

		c.JSON(200, gin.H{"status": "Отправитель удалён"})
	})

	log.Println("Сервер запущен на :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}
