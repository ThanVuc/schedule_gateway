package team_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/team_service"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type sprintClient struct {
	logger       log.Logger
	sprintClient team_service.SprintServiceClient
}

func (sc *sprintClient) CreateSprint(c *gin.Context, req *team_service.CreateSprintRequest) (*team_service.CreateSprintResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := sc.sprintClient.CreateSprint(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
