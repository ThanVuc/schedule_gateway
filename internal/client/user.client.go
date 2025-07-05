package client

import (
	"context"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/proto/user"
)

type userClient struct {
	logger     *loggers.LoggerZap
	userClient user.UserServiceClient
}

func (u *userClient) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	resp, err := u.userClient.GetUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (u *userClient) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoRequest) (*user.UpdateUserInfoResponse, error) {
	resp, err := u.userClient.UpdateUserInfo(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
