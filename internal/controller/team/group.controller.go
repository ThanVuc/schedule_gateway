package team_controller

import (
	"schedule_gateway/global"
	team_client "schedule_gateway/internal/client/team"
	"schedule_gateway/proto/common"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type GroupController struct {
	logger log.Logger
	client team_client.GroupClient
}

func NewGroupController() *GroupController {
	return &GroupController{
		logger: global.Logger,
		client: team_client.NewGroupClient(),
	}
}

func (gc *GroupController) Ping(ctx *gin.Context) {
	resp, err := gc.client.Ping(ctx, &common.EmptyRequest{})
	if err != nil {
		gc.logger.Error("Failed to ping GroupService: ", "")
	}
	ctx.JSON(200, resp)
}
