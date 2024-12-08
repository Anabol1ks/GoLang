package ws

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Позволяем всем подключениям
	},
}

var clients = make(map[*websocket.Conn]bool)

var broadcast = make(chan Message)

type Message struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}

func HandleConnections(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Ошибка обновления соединения:", err)
		return
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Ошибка чтения сообщения: %v\n", err)
			delete(clients, ws)
			break
		}
		// отправка сообщения в канал
		broadcast <- msg
	}
}

func HandleMessage() {
	for {
		// получение сообщения
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Ошибка при отправке: %v\n", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
