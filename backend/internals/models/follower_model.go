package models

import (
	"time"
)

type Follower struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FollowingID uint      `json:"following_id" gorm:"not null;index"`
	FollowerID  uint      `json:"follower_id" gorm:"not null;index"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}