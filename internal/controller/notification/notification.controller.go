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

func (nc *NotificationController) GetNotificationsByRecipientId(c *gin.Context) {
	userId := c.GetString("user_id")
	idReq := &common.IDRequest{
		Id: userId,
	}
	requestId, _ := c.Get("request-id")
	resp, err := nc.client.GetNotificationsByRecipientId(c, idReq)

	if err != nil {
		nc.logger.Error("Failed to get notifications", requestId.(string))
		c.JSON(500, gin.H{"error": "Failed to get notifications"})
		return
	}

	response.Ok(c, "Get notifications successfully", resp)
}

func (nc *NotificationController) MarkNotificationsAsRead(c *gin.Context) {
	idReq := &common.IDsRequest{}
	if err := c.ShouldBindJSON(idReq); err != nil {
		nc.logger.Error("Invalid request body", "")
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	requestId, _ := c.Get("request-id")
	_, err := nc.client.MarkNotificationAsRead(c, idReq)
	if err != nil {
		nc.logger.Error("Failed to mark notification as read", requestId.(string))
		c.JSON(500, gin.H{"error": "Failed to mark notification as read"})
		return
	}
	response.Ok(c, "Mark notification as read successfully", nil)
}

func (nc *NotificationController) DeleteNotificationById(c *gin.Context) {
	notificationId := c.Param("id")
	idReq := &common.IDRequest{
		Id: notificationId,
	}
	requestId, _ := c.Get("request-id")
	_, err := nc.client.DeleteNotificationById(c, idReq)
	if err != nil {
		nc.logger.Error("Failed to delete notification", requestId.(string))
		c.JSON(500, gin.H{"error": "Failed to delete notification"})
		return
	}
	response.Ok(c, "Delete notification successfully", nil)
}
