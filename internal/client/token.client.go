package client

import (
	"context"
	"fmt"
	"schedule_gateway/global"
	v1Token "schedule_gateway/internal/grpc/token.v1"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/settings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ITokenClient interface {
	RefreshToken(refreshToken string) (*v1Token.RefreshTokenResponse, error)
	RevokeToken(token string) (*v1Token.RevokeTokenResponse, error)
}

type TokenClient struct {
	logger      *loggers.LoggerZap
	config      *settings.AuthService
	tokenClient v1Token.TokenServiceClient
}

func NewTokenClient() ITokenClient {
	logger := global.Logger
	config := global.Config.AuthService

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorString("Failed to connect to TokenService", zap.String("host", config.Host), zap.Int("port", config.Port), zap.Error(err))
		return nil
	}

	client := v1Token.NewTokenServiceClient(conn)
	if client == nil {
		logger.ErrorString("Failed to create TokenService client", zap.String("host", config.Host), zap.Int("port", config.Port))
		return nil
	}

	return &TokenClient{
		logger:      logger,
		config:      &config,
		tokenClient: client,
	}
}

func (t *TokenClient) RefreshToken(refreshToken string) (*v1Token.RefreshTokenResponse, error) {
	req := &v1Token.RefreshTokenRequest{
		RefreshToken: refreshToken,
	}
	resp, err := t.tokenClient.RefreshToken(context.Background(), req)
	if err != nil {
		t.logger.ErrorString("RefreshToken failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (t *TokenClient) RevokeToken(token string) (*v1Token.RevokeTokenResponse, error) {
	req := &v1Token.RevokeTokenRequest{
		Token: token,
	}
	resp, err := t.tokenClient.RevokeToken(context.Background(), req)
	if err != nil {
		t.logger.ErrorString("RevokeToken failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
