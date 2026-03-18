package models

import "time"

type React struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	PostID    uint      `json:"post_id" gorm:"not null;index"`
	UserID    uint      `json:"user_id" gorm:"not null;index"`
	Type      string    `json:"type" gorm:"type:varchar(10);default:'like';check:type IN ('like', 'haha', 'love', 'care', 'sad', 'wow', 'angry')"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}