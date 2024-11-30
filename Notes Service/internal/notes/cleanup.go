package notes

import (
	"log"
	"time"

	"notes-service/internal/storage"
)

func CleanupDeletedNotes() {
	// Удаляем записи, удалённые более 15 дней назад
	expirationDate := time.Now().AddDate(0, 0, -15) // N дней назад

	if err := storage.DB.Unscoped().Where("deleted_at < ?", expirationDate).Delete(&Note{}).Error; err != nil {
		log.Printf("Failed to cleanup notes: %v", err)
	} else {
		log.Println("Cleanup completed: old notes permanently deleted.")
	}
}
