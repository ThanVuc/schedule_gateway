package team_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/team_service"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type usertClient struct {
	logger     log.Logger
	userClient team_service.UserServiceClient
}

func (uc *usertClient) GetUserInfo(c *gin.Context, req *common.EmptyRequest) (*team_service.GetUserInfoResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := uc.userClient.GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (uc *usertClient) NotificationConfiguration(c *gin.Context, req *team_service.NotificationConfigurationRequest) (*team_service.NotificationConfigurationResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := uc.userClient.NotificationConfiguration(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
