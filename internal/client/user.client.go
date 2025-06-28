package client

import (
	"context"
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/internal/grpc/user"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/settings"

	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IUserClient interface {
	GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error)
	UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoRequest) (*user.UpdateUserInfoResponse, error)
}

type UserClient struct {
	logger     *loggers.LoggerZap
	config     *settings.UserService
	userClient user.UserServiceClient
}

func NewUserClient() IUserClient {
	logger := global.Logger
	config := global.Config.UserService

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorString("Failed to connect to UserService", zap.String("host", config.Host), zap.Int("port", config.Port), zap.Error(err))
		return nil
	}

	client := user.NewUserServiceClient(conn)
	if client == nil {
		logger.ErrorString("Failed to create UserService client", zap.String("host", config.Host), zap.Int("port", config.Port))
		return nil
	}

	return &UserClient{
		logger:     logger,
		config:     &config,
		userClient: client,
	}
}

func (a *UserClient) GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error) {
	resp, err := a.userClient.GetUserInfo(ctx, req)
	if err != nil {
		a.logger.ErrorString("Failed to get user info", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *UserClient) UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoRequest) (*user.UpdateUserInfoResponse, error) {
	resp, err := a.userClient.UpdateUserInfo(ctx, req)
	if err != nil {
		a.logger.ErrorString("Failed to update user info", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
