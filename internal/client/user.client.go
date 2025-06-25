package client

import (
	"context"
	"fmt"
	"schedule_gateway/global"
	v1 "schedule_gateway/internal/grpc/user.v1"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/settings"

	"go.uber.org/zap"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IUserClient interface {
	GetUserInfo(userId string) (*v1.GetUserInfoResponse, error)
	UpdateUserInfo(userId string, name string, email string) (*v1.UpdateUserInfoResponse, error)
}

type UserClient struct {
	logger     *loggers.LoggerZap
	config     *settings.UserService
	userClient v1.UserServiceClient
}

func NewUserClient() IUserClient {
	logger := global.Logger
	config := global.Config.UserService

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorString("Failed to connect to UserService", zap.String("host", config.Host), zap.Int("port", config.Port), zap.Error(err))
		return nil
	}

	client := v1.NewUserServiceClient(conn)
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

func (a *UserClient) GetUserInfo(userId string) (*v1.GetUserInfoResponse, error) {
	req := &v1.GetUserInfoRequest{
		UserId: userId,
	}
	resp, err := a.userClient.GetUserInfo(context.Background(), req)
	if err != nil {
		a.logger.ErrorString("Failed to get user info", zap.String("userId", userId), zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *UserClient) UpdateUserInfo(userId string, name string, email string) (*v1.UpdateUserInfoResponse, error) {
	req := &v1.UpdateUserInfoRequest{
		UserId: userId,
		Name:   name,
		Email:  email,
	}
	resp, err := a.userClient.UpdateUserInfo(context.Background(), req)
	if err != nil {
		a.logger.ErrorString("Failed to update user info", zap.String("userId", userId), zap.Error(err))
		return nil, err
	}
	return resp, nil
}
