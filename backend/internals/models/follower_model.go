package models

import (
	"time"
)

type Follower struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	FollowingID uint      `json:"following_id" gorm:"not null;index"`
	Following   User      `json:"following" gorm:"foreignKey:FollowingID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	FollowerID  uint      `json:"follower_id" gorm:"not null;index"`
	Follower    User      `json:"follower" gorm:"foreignKey:FollowerID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
}