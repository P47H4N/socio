package models

import (
	"time"
	"gorm.io/gorm"
)

type Post struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	Content   string         `json:"content" gorm:"type:text"`
	MediaURL  string         `json:"media_url" gorm:"type:varchar(255)"`
	PostType  string         `json:"post_type" gorm:"type:varchar(20);default:'profile'"`
	Privacy   string         `json:"privacy" gorm:"type:varchar(20);default:'public'"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"autoUpdateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	User      User           `json:"user" gorm:"foreignKey:UserID"`
}