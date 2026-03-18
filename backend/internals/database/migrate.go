package database

import (
	"github.com/P47H4N/socio/internals/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.BlockList{})
	db.AutoMigrate(&models.Comment{})
	db.AutoMigrate(&models.Follower{})
	db.AutoMigrate(&models.Message{})
	db.AutoMigrate(&models.Notification{})
	db.AutoMigrate(&models.Post{})
	db.AutoMigrate(&models.React{})
	db.AutoMigrate(&models.Report{})
	db.AutoMigrate(&models.SavedPost{})
	db.AutoMigrate(&models.Security{})
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.UserSetting{})
}