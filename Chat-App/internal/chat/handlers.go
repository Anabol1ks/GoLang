package chat

import (
	"chat_app/internal/storage"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoomMessage(c *gin.Context) {
	roomID := c.Param("room_id")

	var messages []Message
	if err := storage.DB.Where("room_id = ?", roomID).Order("created_at").Find(&messages).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить сообщение"})
		return
	}

	c.JSON(http.StatusOK, messages)
}

func GetRoomsHandler(c *gin.Context) {
	var rooms []Room
	if err := storage.DB.Find(&rooms).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить список комнат"})
		return
	}

	c.JSON(http.StatusOK, rooms)
}

func CreateRoomHandler(c *gin.Context) {
	var input struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Название комнаты обязательно"})
		return
	}

	var existingRoom Room
	if err := storage.DB.Where("name = ?", input.Name).First(&existingRoom).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Комната с таким названием уже существует"})
		return
	}

	room := Room{Name: input.Name}
	if err := storage.DB.Create(&room).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось создать комнату"})
		return
	}

	c.JSON(http.StatusCreated, room)
}
