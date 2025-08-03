package auth_client

import (
	"schedule_gateway/proto/auth"

	"github.com/thanvuc/go-core-lib/log"
)

type userClient struct {
	logger     log.Logger
	userClient auth.UserServiceClient
}
