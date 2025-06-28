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

type IAuthClient interface {
	Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error)
	Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error)
	ConfirmEmail(ctx context.Context, req *auth.ConfirmEmailRequest) (*auth.ConfirmEmailResponse, error)
	Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error)
	ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error)
	ForgotPassword(ctx context.Context, req *auth.ForgotPasswordRequest) (*auth.ForgotPasswordResponse, error)
	ConfirmForgotPassword(ctx context.Context, req *auth.ConfirmForgotPasswordRequest) (*auth.ConfirmForgotPasswordResponse, error)
	SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error)
}

type AuthClient struct {
	logger     *loggers.LoggerZap
	config     *settings.AuthService
	authClient auth.AuthServiceClient
}

func NewAuthClient() IAuthClient {
	logger := global.Logger
	config := global.Config.AuthService

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorString("Failed to connect to AuthService", zap.String("host", config.Host), zap.Int("port", config.Port), zap.Error(err))
		return nil
	}

	client := auth.NewAuthServiceClient(conn)
	if client == nil {
		logger.ErrorString("Failed to create AuthService client", zap.String("host", config.Host), zap.Int("port", config.Port))
		return nil
	}

	return &AuthClient{
		logger:     logger,
		config:     &config,
		authClient: client,
	}
}

func (a *AuthClient) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	resp, err := a.authClient.Login(ctx, req)
	if err != nil {
		a.logger.ErrorString("Login failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	resp, err := a.authClient.Register(ctx, req)
	if err != nil {
		a.logger.ErrorString("Register failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) ConfirmEmail(ctx context.Context, req *auth.ConfirmEmailRequest) (*auth.ConfirmEmailResponse, error) {
	resp, err := a.authClient.ConfirmEmail(ctx, req)
	if err != nil {
		a.logger.ErrorString("ConfirmEmail failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	resp, err := a.authClient.Logout(ctx, req)
	if err != nil {
		a.logger.ErrorString("Logout failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	resp, err := a.authClient.ResetPassword(ctx, req)
	if err != nil {
		a.logger.ErrorString("ResetPassword failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) ForgotPassword(ctx context.Context, req *auth.ForgotPasswordRequest) (*auth.ForgotPasswordResponse, error) {
	resp, err := a.authClient.ForgotPassword(ctx, req)
	if err != nil {
		a.logger.ErrorString("ForgotPassword failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) ConfirmForgotPassword(ctx context.Context, req *auth.ConfirmForgotPasswordRequest) (*auth.ConfirmForgotPasswordResponse, error) {
	resp, err := a.authClient.ConfirmForgotPassword(ctx, req)
	if err != nil {
		a.logger.ErrorString("ConfirmForgotPassword failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error) {
	resp, err := a.authClient.SaveRouteResource(ctx, req)
	if err != nil {
		a.logger.ErrorString("SaveRouteResource failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
