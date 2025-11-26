package personalschedule_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/personal_schedule"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type workClient struct {
	logger     log.Logger
	workClient personal_schedule.WorkServiceClient
}

func (wc *workClient) UpsertWork(c *gin.Context, req *personal_schedule.UpsertWorkRequest) (*personal_schedule.UpsertWorkResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))
	resp, err := wc.workClient.UpsertWork(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) GetWorks(c *gin.Context, req *personal_schedule.GetWorksRequest) (*personal_schedule.GetWorksResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))
	resp, err := wc.workClient.GetWorks(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) GetWork(c *gin.Context, req *personal_schedule.GetWorkRequest) (*personal_schedule.GetWorkResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))
	resp, err := wc.workClient.GetWork(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (wc *workClient) DeleteWork(c *gin.Context, req *personal_schedule.DeleteWorkRequest) (*personal_schedule.DeleteWorkResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))
	resp, err := wc.workClient.DeleteWork(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}