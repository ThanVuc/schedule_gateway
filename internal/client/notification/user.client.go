package notification_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/notification_service"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type userNotificationClient struct {
	logger                 log.Logger
	userNotificationClient notification_service.UserNotificationServiceClient
}

func (uc *userNotificationClient) UpsertUserFCMToken(c *gin.Context, req *notification_service.UpsertUserFCMTokenRequest) (*common.EmptyResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := uc.userNotificationClient.UpsertUserFCMToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
