package setting

type SettingsBody struct {
	IsPrivateAccount  *bool   `json:"is_private_account"`
	AllowMessage      *string `json:"allow_message" binding:"omitempty,oneof=everyone none friends followers"`
	EmailNotification *bool   `json:"email_notification"`
	PushNotification  *bool   `json:"push_notification"`
	Language          *string `json:"language"`
}