package ws

import (
	"chat_app/internal/auth"
	"chat_app/internal/chat"
	"chat_app/internal/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type RoomClients struct {
	Clients   map[*websocket.Conn]bool
	Broadcast chan Message
}

var rooms = make(map[uint]*RoomClients)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Позволяем всем подключениям
	},
}

type Message struct {
	RoomID   uint   `json:"room_id"`
	Username string `json:"username"`
	Content  string `json:"content"`
}

func HandleConnections(c *gin.Context) {
	userID := c.GetUint("user_id")

	var user auth.User
	if err := storage.DB.First(&user, userID).Error; err != nil {
		log.Println("Пользователь не найден")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось найти пользователя"})
		return
	}

	roomID, err := strconv.Atoi(c.Param("room_id"))
	if err != nil {
		log.Println("Неверный id комнаты")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный id комнаты"})
		return
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Не удалось обновить соединение:", err)
		return
	}
	defer func() {
		room := rooms[uint(roomID)]
		delete(room.Clients, ws)
		ws.Close()

		room.Broadcast <- Message{
			RoomID:   uint(roomID),
			Username: user.Username,
			Content:  "Покинул комнату",
		}
	}()

	if rooms[uint(roomID)] == nil {
		rooms[uint(roomID)] = &RoomClients{
			Clients:   make(map[*websocket.Conn]bool),
			Broadcast: make(chan Message),
		}
		go HandleRoomMessages(uint(roomID))
	}

	room := rooms[uint(roomID)]
	room.Clients[ws] = true

	room.Broadcast <- Message{
		RoomID:   uint(roomID),
		Username: user.Username,
		Content:  "Присоединился к комнате",
	}

	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Ошибка чтения сообщения: %v\n", err)
			delete(room.Clients, ws)
			break
		}

		msg.RoomID = uint(roomID)
		room.Broadcast <- msg

		msgDB := chat.Message{
			RoomID:  msg.RoomID,
			UserID:  userID,
			Content: msg.Content,
		}
		if err := storage.DB.Create(&msgDB).Error; err != nil {
			log.Printf("Ошибка сохранения сообщение: %v\n", err)
		}
	}

}

func HandleRoomMessages(roomID uint) {
	room := rooms[roomID]
	for {
		// Получаем сообщение из канала комнаты
		msg := <-room.Broadcast

		// Отправляем сообщение всем клиентам в комнате
		for client := range room.Clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("Ошибка при отправке: %v\n", err)
				client.Close()
				delete(room.Clients, client)
			}
		}
	}
}
