package chat

import "gorm.io/gorm"

type Room struct {
	Name       string `gorm:"type:varchar(100);unique;not null"`
	gorm.Model `swaggerignore:"true"`
}

type Message struct {
	gorm.Model
	RoomID  uint   `gorm:"not null"`
	UserID  uint   `gorm:"not null"`
	Content string `gorm:"type:text;not null"`
}
