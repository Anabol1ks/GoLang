package notes

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title       string `gorm:"type:varchar(100);not null"`
	Description string `gorm:"type:text"`
	UserID      uint   `gorm:"not null"`
}
