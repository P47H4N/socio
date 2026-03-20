package setting

import (
	"errors"

	"github.com/P47H4N/socio/internals/models"
	"gorm.io/gorm"
)

type SettingService struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *SettingService{
	return &SettingService{
		db: db,
	}
}

func (ss *SettingService) GetSetting(uid uint) (*models.UserSetting, error) {
	var userSetting models.UserSetting
	if err := ss.db.Where("user_id = ?", uid).First(&userSetting); err != nil {
		return nil, errors.New("No user settings found.")
	}
	return &userSetting, nil
}

func (ss *SettingService) UpdateSetting(body *SettingsBody, uid uint) error {
	updateSetting := make(map[string]interface{})
	if body.IsPrivateAccount != nil {
        updateSetting["is_private_account"] = *body.IsPrivateAccount
    } else if body.AllowMessage != nil {
        updateSetting["allow_message"] = *body.AllowMessage
    } else if body.EmailNotification != nil {
        updateSetting["email_notification"] = *body.EmailNotification
    } else if body.PushNotification != nil {
        updateSetting["push_notification"] = *body.PushNotification
    } else if body.Language != nil {
        updateSetting["language"] = *body.Language
    }
	if len(updateSetting) == 0 {
        return errors.New("No fields to update.")
    }
    if err := ss.db.Model(&models.UserSetting{}).Where("user_id = ?", uid).Updates(updateSetting).Error; err != nil {
        return errors.New("Failed to update data.")
    }
	return nil
}
