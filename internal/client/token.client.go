package client

import (
	"context"
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/internal/grpc/auth"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/settings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ITokenClient interface {
	RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error)
	RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error)
}

type TokenClient struct {
	logger      *loggers.LoggerZap
	config      *settings.AuthService
	tokenClient auth.TokenServiceClient
}

func NewTokenClient() ITokenClient {
	logger := global.Logger
	config := global.Config.AuthService

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorString("Failed to connect to TokenService", zap.String("host", config.Host), zap.Int("port", config.Port), zap.Error(err))
		return nil
	}

	client := auth.NewTokenServiceClient(conn)
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

func (t *TokenClient) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	resp, err := t.tokenClient.RefreshToken(ctx, req)
	if err != nil {
		t.logger.ErrorString("RefreshToken failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (t *TokenClient) RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error) {
	resp, err := t.tokenClient.RevokeToken(ctx, req)
	if err != nil {
		t.logger.ErrorString("RevokeToken failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
