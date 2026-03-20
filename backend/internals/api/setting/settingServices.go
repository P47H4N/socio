package setting

import (
	"errors"
	"time"

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

func (ss *SettingService) GetUserReports(uid uint) ([]models.Report, error) {
	var reports []models.Report
	if err := ss.db.Where("reporter_id = ?", uid).Find(&reports).Error; err != nil {
		return nil, errors.New("Unable to fetch data.")
	}
	if len(reports) == 0 {
		return nil, errors.New("No reports found.")
	}
	return reports, nil
}

func (ss *SettingService) GetReportDetails(rid, uid uint) (*models.Report, error) {
	var report models.Report
	if err := ss.db.Where("id = ? AND reporter_id = ?", rid, uid).First(&report).Error; err != nil {
		return nil, errors.New("Unable to fetch report details.")
	}
	return &report, nil
}

func (ss *SettingService) SubmitReport(uid uint, body *ReportBody) error {
	report := models.Report{
		ReporterID: uid,
		TargetType: body.TargetType,
		TargetID: body.TargetID,
		Reason: body.Reason,
		CreatedAt: time.Now(),
	}
	if err := ss.db.Create(&report).Error; err != nil {
		return errors.New("Report create failed.")
	}
	return nil
}