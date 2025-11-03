package personalschedule_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/personal_schedule"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type goalClient struct {
	logger     log.Logger
	goalClient personal_schedule.GoalServiceClient
}

func (gc *goalClient) GetGoals(c *gin.Context, req *personal_schedule.GetGoalsRequest) (*personal_schedule.GetGoalsResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))
	resp, err := gc.goalClient.GetGoals(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
