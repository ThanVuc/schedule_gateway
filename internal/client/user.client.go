package client

import (
	"context"
	"schedule_gateway/internal/grpc/user"
	"schedule_gateway/pkg/loggers"

	"go.uber.org/zap"
)

type userClient struct {
	logger     *loggers.LoggerZap
	userClient user.UserServiceClient
}

func (u *userClient) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	resp, err := u.userClient.GetUserInfo(ctx, req)
	if err != nil {
		u.logger.ErrorString("Failed to get user info", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (u *userClient) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoRequest) (*user.UpdateUserInfoResponse, error) {
	resp, err := u.userClient.UpdateUserInfo(ctx, req)
	if err != nil {
		u.logger.ErrorString("Failed to update user info", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
