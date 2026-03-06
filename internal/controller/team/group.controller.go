package team_controller

import (
	"schedule_gateway/global"
	team_client "schedule_gateway/internal/client/team"

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
