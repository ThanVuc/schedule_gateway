package notification_controller

import (
	"schedule_gateway/global"
	notification_client "schedule_gateway/internal/client/notification"
	"schedule_gateway/pkg/response"
	"schedule_gateway/proto/common"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type NotificationController struct {
	client notification_client.NotificationClient
	logger log.Logger
}

func NewNotificationController() *NotificationController {
	client := notification_client.NewNotificationClient()
	logger := global.Logger
	return &NotificationController{
		client: client,
		logger: logger,
	}
}

func (nc *NotificationController) GetNotifications(ctx *gin.Context) {
	idReq := common.IDRequest{
		Id: "Hello World",
	}
	requestId, _ := ctx.Get("request-id")
	resp, err := nc.client.GetNotifications(ctx, &idReq)

	if err != nil {
		nc.logger.Error("Failed to get notifications", requestId.(string))
		ctx.JSON(500, gin.H{"error": "Failed to get notifications"})
		return
	}

	response.Ok(ctx, "Get notifications successfully", resp)
}
