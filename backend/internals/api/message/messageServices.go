package message

import "gorm.io/gorm"

type MessageService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *MessageService {
	return &MessageService{
		db: db,
	}
}

