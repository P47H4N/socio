package models

import (
	"time"
)

type Report struct {
	ID         uint      `json:"-" gorm:"primaryKey"`
	ReporterID uint      `json:"reporter_id" gorm:"not null;index"`
	TargetType string    `json:"target_type" gorm:"type:varchar(20);not null;check:target_type IN ('profile', 'post', 'comment')"`
	TargetID   uint      `json:"target_id" gorm:"not null"`
	Reason     string    `json:"reason" gorm:"type:text;not null"`
	Status     string    `json:"status" gorm:"type:varchar(20);default:'Submitted';check:status IN ('Submitted', 'Reviewed', 'Rejected', 'Actioned')"`
	CreatedAt  time.Time `json:"created_at" gorm:"autoCreateTime"`
}