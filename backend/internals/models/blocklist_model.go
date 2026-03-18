package models

import (
	"time"
)

type BlockList struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	BlockerID uint      `json:"blocker_id" gorm:"not null;index"`
	BlockedID uint      `json:"blocked_id" gorm:"not null;index"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
