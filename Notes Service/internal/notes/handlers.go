package notes

import (
	"net/http"
	"notes-service/internal/storage"

	"github.com/gin-gonic/gin"
)

func CreateNoteHandler(c *gin.Context) {
	userID := c.GetUint("user_id")

	var note Note
	note.UserID = userID
	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := storage.DB.Create(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка создания заметки"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func GetNotesHandler(c *gin.Context) {
	userID := c.GetUint("user_id") // Получаем ID текущего пользователя

	search := c.Query("search")   //поиск по названию
	startDate := c.Query("start") // Начальная карта
	endDate := c.Query("end")     // Конечная карта

	var notes []Note
	query := storage.DB.Where("user_id = ?", userID)

	if search != "" {
		query = query.Where("title ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
	}

	if startDate != "" && endDate != "" {
		query = query.Where("created_at BETWEEN ? AND ?", startDate, endDate)
	}

	if err := query.Find(&notes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notes"})
		return
	}

	c.JSON(http.StatusOK, notes)
}

func GetNoteHandler(c *gin.Context) {
	id := c.Param("id")
	var note Note
	if err := storage.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заметка не найдена"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func UpdateNoteHandler(c *gin.Context) {
	id := c.Param("id")
	var note Note
	if err := storage.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заметка не найдена"})
		return
	}
	userID := c.GetUint("user_id")
	if note.UserID != userID {
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		return
	}

	if err := c.ShouldBindJSON(&note); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := storage.DB.Save(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка обновления заметок"})
		return
	}

	c.JSON(http.StatusOK, note)
}

func DeleteNoteHandler(c *gin.Context) {
	userID := c.GetUint("user_id")
	id := c.Param("id")

	var note Note
	if err := storage.DB.Where("id = ? AND user_id = ?", id, userID).First(&note).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Note not found"})
		return
	}

	if err := storage.DB.Delete(&note).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete note"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Note deleted"})
}
