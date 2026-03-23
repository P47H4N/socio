package models

import (
	"time"

	"gorm.io/gorm"
)

type Message struct {
	ID         uint           `json:"id" gorm:"primaryKey"`
	SenderID   uint           `json:"sender_id" gorm:"not null;index"`
	ReceiverID uint           `json:"receiver_id" gorm:"not null;index"`
	Message    string         `json:"message" gorm:"type:text"`
	MediaUrl   string         `json:"media_url" gorm:"type:varchar(255)"`
	IsRead     string         `json:"is_read" gorm:"type:varchar(15);default:'sent';check:is_read IN ('sent', 'delivered', 'seen')"`
	CreatedAt  time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
