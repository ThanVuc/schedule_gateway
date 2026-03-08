package team_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/team_service"

	"schedule_gateway/proto/common"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type groupClient struct {
	logger      log.Logger
	groupClient team_service.GroupServiceClient
}

func (wc *groupClient) Ping(c *gin.Context, req *common.EmptyRequest) (*common.EmptyResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.groupClient.Ping(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
