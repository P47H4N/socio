package models

import (
	"time"

	"gorm.io/gorm"
)

type Comment struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	PostID    uint           `json:"post_id" gorm:"not null;index"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	ParentID  *uint          `json:"parent_id" gorm:"index"`
	Content   string         `json:"content" gorm:"type:text"`
	MediaURL  string         `json:"media_url" gorm:"type:varchar(255)"`
	CreatedAt time.Time      `json:"created_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
	Replies   []Comment      `json:"replies,omitempty" gorm:"foreignKey:ParentID"`
	User      User           `json:"user" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
