package personalschedule_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/personal_schedule"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type labelClient struct {
	logger      log.Logger
	labelClient personal_schedule.LabelServiceClient
}

func (lc *labelClient) GetLabelPerTypes(c *gin.Context, req *common.EmptyRequest) (*personal_schedule.GetLabelPerTypesResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := lc.labelClient.GetLabelPerTypes(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
