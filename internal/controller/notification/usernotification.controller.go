package notification_controller

import (
	"schedule_gateway/global"
	notification_client "schedule_gateway/internal/client/notification"
	dtos "schedule_gateway/internal/dtos/notification"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/notification_service"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
	"go.uber.org/zap"
)

type UserNotificationController struct {
	client notification_client.UserNotificationClient
	logger log.Logger
}

func NewUserNotificationController() *UserNotificationController {
	client := notification_client.NewUserNotificationClient()
	logger := global.Logger
	return &UserNotificationController{
		client: client,
		logger: logger,
	}
}

func (uc *UserNotificationController) UpsertUserFCMToken(c *gin.Context) {
	requestDto := dtos.UpsertNotificationRequestDTO{}

	requestId := c.GetString("request-id")
	userId := c.GetString("user_id")

	if err := c.ShouldBindJSON(&requestDto); err != nil {
		response.BadRequest(c, "Bad Request Error")
		return
	}
	request := &notification_service.UpsertUserFCMTokenRequest{
		UserId:   userId,
		FcmToken: requestDto.FCMToken,
		DeviceId: requestDto.DeviceID,
	}
	resp, err := uc.client.UpsertUserFCMToken(c, request)
	if err != nil {
		uc.logger.Error("Failed to upsert user FCM token", requestId, zap.Error(err))
		response.InternalServerError(c, "Internal Server Error")
		return
	}

	response.Ok(c, "Call API Successfully", resp)
}
