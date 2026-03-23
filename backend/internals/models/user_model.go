package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID             uint           `json:"id,omitempty" gorm:"primaryKey"`
	Username       string         `json:"username,omitempty" gorm:"type:varchar(50);uniqueIndex;not null"`
	Email          string         `json:"email,omitempty" gorm:"type:varchar(100);uniqueIndex;not null"`
	Phone          *string        `json:"phone,omitempty" gorm:"type:varchar(20);uniqueIndex;default:null"`
	FullName       string         `json:"full_name,omitempty" gorm:"type:varchar(100);not null"`
	Password       string         `json:"-" gorm:"type:varchar(255);not null"`
	Bio            string         `json:"bio,omitempty" gorm:"type:text"`
	ProfilePic     string         `json:"profile_pic,omitempty" gorm:"type:varchar(255)"`
	CoverPic       string         `json:"cover_pic,omitempty" gorm:"type:varchar(255)"`
	LastSeen       time.Time      `json:"last_seen,omitempty" gorm:"autoUpdateTime"`
	AccountStatus  string         `json:"account_status,omitempty" gorm:"type:varchar(20);default:'active'"`
	CreatedAt      time.Time      `json:"created_at,omitempty" gorm:"autoCreateTime"`
	UpdatedAt      time.Time      `json:"-" gorm:"autoUpdateTime"`
	DeletedAt      gorm.DeletedAt `json:"-" gorm:"index"`
	FollowerCount  int64          `json:"follower_count" gorm:"-"`
	FollowingCount int64          `json:"following_count" gorm:"-"`
}
