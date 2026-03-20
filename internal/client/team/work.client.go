package team_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/team_service"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type workClient struct {
	logger     log.Logger
	workClient team_service.WorkServiceClient
}

func (wc *workClient) CreateWork(c *gin.Context, req *team_service.CreateWorkRequest) (*team_service.CreateWorkResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.CreateWork(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) GetWork(c *gin.Context, req *common.IDRequest) (*team_service.GetWorkResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.GetWork(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) ListWorks(c *gin.Context, req *team_service.ListWorksRequest) (*team_service.ListWorksResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.ListWorks(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) UpdateWork(c *gin.Context, req *team_service.UpdateWorkRequest) (*team_service.UpdateWorkResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.UpdateWork(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) DeleteWork(c *gin.Context, req *common.IDRequest) (*team_service.DeleteWorkResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.DeleteWork(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) CreateChecklistItem(c *gin.Context, req *team_service.CreateChecklistItemRequest) (*team_service.CreateChecklistItemResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.CreateChecklistItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) UpdateChecklistItem(c *gin.Context, req *team_service.UpdateChecklistItemRequest) (*team_service.UpdateChecklistItemResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.UpdateChecklistItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) DeleteChecklistItem(c *gin.Context, req *common.IDRequest) (*team_service.DeleteChecklistItemResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.DeleteChecklistItem(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) CreateComment(c *gin.Context, req *team_service.CreateCommentRequest) (*team_service.CreateCommentResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.CreateComment(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) UpdateComment(c *gin.Context, req *team_service.UpdateCommentRequest) (*team_service.UpdateCommentResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.UpdateComment(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) DeleteComment(c *gin.Context, req *common.IDRequest) (*team_service.DeleteCommentResponse, error) {
	ctx := context.Background()
	ctx = utils.EnrichContext(ctx, c)
	resp, err := wc.workClient.DeleteComment(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
