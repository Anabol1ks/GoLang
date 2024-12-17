package ws

import (
	"chat_app/internal/auth"
	"chat_app/internal/chat"
	"chat_app/internal/storage"
	"fmt"
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
	Content string `json:"content"`
}

// HandleConnections godoc
// @Summary Подключение к комнате
// @Description Устанавливает WebSocket-соединение с указанной комнатой. Используйте ws:// или wss:// для подключения.
// @Tags ws
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param input body Message true "Сообщение"
// @Param room_id path int true "ID комнаты"
// @Failure 400 {object} swg.ErrorResponse "Неверный ID комнаты"
// @Failure 409 {object} swg.ErrorResponse "Не удалось найти пользователя"
// @Router /ws/{room_id} [get]
func HandleConnections(c *gin.Context) {
	// Извлекаем user_id из контекста
	userID := c.GetUint("user_id")

	// Получаем данные пользователя
	var user auth.User
	if err := storage.DB.First(&user, userID).Error; err != nil {
		log.Println("Пользователь не найден")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Не удалось найти пользователя"})
		return
	}

	// Получаем ID комнаты
	roomID, err := strconv.Atoi(c.Param("room_id"))
	if err != nil {
		log.Println("Неверный ID комнаты")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID комнаты"})
		return
	}

	// Обновляем соединение до WebSocket
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("Не удалось обновить соединение:", err)
		return
	}
	defer func() {
		room := rooms[uint(roomID)]
		delete(room.Clients, ws)
		ws.Close()

		// Уведомляем об уходе пользователя
		room.Broadcast <- Message{
			Content: fmt.Sprintf("%s покинул комнату", user.Username),
		}
	}()

	// Регистрируем комнату, если её ещё нет
	if rooms[uint(roomID)] == nil {
		rooms[uint(roomID)] = &RoomClients{
			Clients:   make(map[*websocket.Conn]bool),
			Broadcast: make(chan Message),
		}
		go HandleRoomMessages(uint(roomID))
	}

	// Регистрируем клиента в комнате
	room := rooms[uint(roomID)]
	room.Clients[ws] = true

	// Уведомляем о присоединении
	room.Broadcast <- Message{
		Content: fmt.Sprintf("%s присоединился к комнате", user.Username),
	}

	// Чтение сообщений
	for {
		var receivedMsg Message
		err := ws.ReadJSON(&receivedMsg)
		if err != nil {
			log.Printf("Ошибка чтения сообщения: %v\n", err)
			break
		}

		// Добавляем имя пользователя к сообщению
		msg := Message{
			Content: fmt.Sprintf("%s: %s", user.Username, receivedMsg.Content),
		}
		room.Broadcast <- msg

		// Сохраняем сообщение в базе данных
		msgDB := chat.Message{
			RoomID:  uint(roomID),
			UserID:  userID,
			Content: receivedMsg.Content,
		}
		if err := storage.DB.Create(&msgDB).Error; err != nil {
			log.Printf("Ошибка сохранения сообщения: %v\n", err)
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
