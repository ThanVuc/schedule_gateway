package client

import (
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/proto/user"
)

type userClient struct {
	logger     *loggers.LoggerZap
	userClient user.UserServiceClient
}
