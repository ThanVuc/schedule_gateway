package dtos

type TeamUserDTO struct {
	Email                string `json:"email"`
	UseEmailNotification bool
	UseAppNotification   bool
	CreateAt             string `json:"created_at"`
}

type NotificationConfigurationRequestDTO struct {
	UseEmailNotification bool  `json:"use_email_notification"`
	UseAppNotification   *bool `json:"use_app_notification"`
}
