package models

import (
	"time"
)

type Security struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"type:varchar(100);index;not null"`
	Token     string    `json:"token" gorm:"type:varchar(255);not null"`
	Type      string    `json:"type" gorm:"type:varchar(30);not null;check:type IN ('email verification', 'password reset')"`
	IsUsed    bool      `json:"is_used" gorm:"default:false"`
	ExpiredAt time.Time `json:"expired_at" gorm:"not null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
}
