package team_client

import (
	"schedule_gateway/proto/team_service"

	"github.com/thanvuc/go-core-lib/log"
)

type sprintClient struct {
	logger     log.Logger
	teamClient team_service.SprintServiceClient
}
