package notification_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/notification_service"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type notificationClient struct {
	logger             log.Logger
	notificationClient notification_service.NotificationServiceClient
}

func (nc *notificationClient) GetNotificationsByRecipientId(c *gin.Context, req *common.IDRequest) (*notification_service.GetNotificationsByRecipientIdResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := nc.notificationClient.GetNotificationsByRecipientId(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (nc *notificationClient) MarkNotificationAsRead(c *gin.Context, req *common.IDsRequest) (*common.EmptyResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := nc.notificationClient.MarkNotificationsAsRead(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (nc *notificationClient) DeleteNotificationById(c *gin.Context, req *common.IDRequest) (*common.EmptyResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := nc.notificationClient.DeleteNotificationById(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
