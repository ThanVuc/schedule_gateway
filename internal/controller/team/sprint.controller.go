package team_controller

import (
	"schedule_gateway/global"
	team_client "schedule_gateway/internal/client/team"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type SprintController struct {
	logger log.Logger
	client team_client.SprintClient
}

func NewSprintController() *SprintController {
	return &SprintController{
		logger: global.Logger,
		client: team_client.NewSprintClient(),
	}
}


func (sc *SprintController) CreateSprint(ctx *gin.Context) {
	
}