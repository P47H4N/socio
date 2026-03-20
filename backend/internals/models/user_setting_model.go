package models

type UserSetting struct {
	ID                uint   `json:"-" gorm:"primaryKey"`
	UserID            uint   `json:"user_id" gorm:"uniqueIndex;not null"`
	IsPrivateAccount  bool   `json:"is_private_account" gorm:"default:false"`
	AllowMessage      string `json:"allow_message" gorm:"type:varchar(20);default:'everyone'"`
	EmailNotification bool   `json:"email_notification" gorm:"default:true"`
	PushNotification  bool   `json:"push_notification" gorm:"default:true"`
	Language          string `json:"language" gorm:"type:varchar(10);default:'en'"`
}