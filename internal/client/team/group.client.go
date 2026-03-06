package team_client

import (
	"schedule_gateway/proto/team_service"

	"github.com/thanvuc/go-core-lib/log"
)

type groupClient struct {
	logger     log.Logger
	teamClient team_service.GroupServiceClient
}
