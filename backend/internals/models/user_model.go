package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            uint           `json:"id" gorm:"primaryKey"`
	Username      string         `json:"username" gorm:"type:varchar(50);uniqueIndex;not null"`
	Email         string         `json:"email" gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone         *string        `json:"phone" gorm:"type:varchar(20);uniqueIndex;default:null"`
	FullName      string         `json:"full_name" gorm:"type:varchar(100);not null"`
	Password      string         `json:"-" gorm:"type:varchar(255);not null"`
	Bio           string         `json:"bio" gorm:"type:text"`
	ProfilePic    string         `json:"profile_pic" gorm:"type:varchar(255)"`
	CoverPic      string         `json:"cover_pic" gorm:"type:varchar(255)"`
	LastSeen      time.Time      `json:"last_seen" gorm:"autoUpdateTime"`
	AccountStatus string         `json:"account_status" gorm:"type:varchar(20);default:'active'"`
	CreatedAt     time.Time      `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt     time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt     gorm.DeletedAt `json:"-" gorm:"index"`
}
