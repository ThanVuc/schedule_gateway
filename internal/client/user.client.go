package client

import (
	"schedule_gateway/proto/user"

	"github.com/thanvuc/go-core-lib/log"
)

type userClient struct {
	logger     log.Logger
	userClient user.UserServiceClient
}
