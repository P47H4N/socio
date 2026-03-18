package models

import (
	"time"

	"gorm.io/gorm"
)

type Notification struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	ActorID   uint           `json:"actor_id" gorm:"not null;index"`
	Type      string         `json:"type" gorm:"type:varchar(20);not null;check:type IN ('post', 'post like', 'post comment', 'reply comment', 'share post', 'follow')"`
	TargetID  uint           `json:"target_id" gorm:"not null"`
	IsRead    bool           `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
