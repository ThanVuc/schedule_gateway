package dtos

type UpsertNotificationRequestDTO struct {
	FCMToken string `json:"fcm_token"`
	DeviceID string `json:"device_id"`
}
