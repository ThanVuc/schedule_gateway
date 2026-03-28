package team_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/common"
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

func (sc *sprintClient) GetSprint(c *gin.Context, req *common.IDRequest) (*team_service.GetSprintResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := sc.sprintClient.GetSprint(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (sc *sprintClient) GetSimpleSprints(c *gin.Context, req *common.IDRequest) (*team_service.GetSimpleSprintsResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := sc.sprintClient.GetSimpleSprints(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (sc *sprintClient) ExportSprint(c *gin.Context, req *common.IDRequest) (*team_service.ExportSprintResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := sc.sprintClient.ExportSprint(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (sc *sprintClient) ListSprints(c *gin.Context, req *team_service.ListSprintsRequest) (*team_service.ListSprintsResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := sc.sprintClient.ListSprints(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (sc *sprintClient) UpdateSprint(c *gin.Context, req *team_service.UpdateSprintRequest) (*team_service.UpdateSprintResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := sc.sprintClient.UpdateSprint(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (sc *sprintClient) UpdateSprintStatus(c *gin.Context, req *team_service.UpdateSprintStatusRequest) (*team_service.UpdateSprintStatusResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := sc.sprintClient.UpdateSprintStatus(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (sc *sprintClient) DeleteSprint(c *gin.Context, req *common.IDRequest) (*team_service.DeleteSprintResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := sc.sprintClient.DeleteSprint(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}
