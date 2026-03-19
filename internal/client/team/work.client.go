package team_client

import (
	"schedule_gateway/proto/team_service"

	"github.com/thanvuc/go-core-lib/log"
)

type workClient struct {
	logger     log.Logger
	workClient team_service.WorkServiceClient
}
