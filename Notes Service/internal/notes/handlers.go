package notes

import (
	"net/http"
	"notes-service/internal/storage"

	"github.com/gin-gonic/gin"
)

// CreateNoteHandler godoc
// @Summary Создание новой заметки
// @Description Создает новую заметку для текущего пользователя
// @Tags notes
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param note body notes.Note true "Данные заметки"
// @Success 200 {object} notes.Note "Создана новая заметка"
// @Failure 400 {object} map[string]string "Ошибка валидации данных"
// @Failure 401 {object} map[string]string "Неавторизованный запрос"
// @Router /notes [post]
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

// GetNotesHandler godoc
// @Summary Получение списка заметок
// @Description Возвращает все заметки текущего пользователя с возможностью фильтрации
// @Tags notes
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param search query string false "Поиск по названию или описанию"
// @Param start query string false "Начальная дата в формате YYYY-MM-DD"
// @Param end query string false "Конечная дата в формате YYYY-MM-DD"
// @Success 200 {array} notes.Note "Список заметок"
// @Failure 401 {object} map[string]string "Неавторизованный запрос"
// @Router /notes [get]
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

// GetNoteHandler godoc
// @Summary Получение заметки по ID
// @Description Возвращает заметку по ID для текущего пользователя
// @Tags notes
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID заметки"
// @Success 200 {object} notes.Note "Заметка"
// @Failure 404 {object} map[string]string "Заметка не найдена"
// @Failure 401 {object} map[string]string "Неавторизованный запрос"
// @Router /notes/{id} [get]
func GetNoteHandler(c *gin.Context) {
	id := c.Param("id")
	var note Note
	if err := storage.DB.First(&note, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Заметка не найдена"})
		return
	}

	c.JSON(http.StatusOK, note)
}

// UpdateNoteHandler godoc
// @Summary Обновление заметки по ID
// @Description Обновляет данные заметки для текущего пользователя
// @Tags notes
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID заметки"
// @Param note body notes.Note true "Данные для обновления"
// @Success 200 {object} notes.Note "Обновленная заметка"
// @Failure 400 {object} map[string]string "Ошибка валидации данных"
// @Failure 401 {object} map[string]string "Неавторизованный запрос"
// @Failure 403 {object} map[string]string "Нет доступа"
// @Failure 404 {object} map[string]string "Заметка не найдена"
// @Router /notes/{id} [put]
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

// DeleteNoteHandler godoc
// @Summary Удаление заметки по ID
// @Description Удаляет заметку для текущего пользователя
// @Tags notes
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "ID заметки"
// @Success 200 {object} map[string]string "Сообщение об удалении"
// @Failure 401 {object} map[string]string "Неавторизованный запрос"
// @Failure 404 {object} map[string]string "Заметка не найдена"
// @Router /notes/{id} [delete]
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
