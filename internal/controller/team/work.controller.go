package team_controller

import (
	"schedule_gateway/global"
	team_client "schedule_gateway/internal/client/team"

	"github.com/thanvuc/go-core-lib/log"
)

type WorkController struct {
	logger log.Logger
	client team_client.WorkClient
}

func NewWorkController() *WorkController {
	return &WorkController{
		logger: global.Logger,
		client: team_client.NewWorkClient(),
	}
}
