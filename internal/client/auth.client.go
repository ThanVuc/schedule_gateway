package client

import (
	"context"
	"fmt"
	"schedule_gateway/global"
	v1Auth "schedule_gateway/internal/grpc/auth.v1"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/settings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IAuthClient interface {
	Login(username, password string) (*v1Auth.LoginResponse, error)
	Register(email, password string) (*v1Auth.RegisterResponse, error)
	ConfirmEmail(email, token, password string) (*v1Auth.ConfirmEmailResponse, error)
	Logout(accessToken, refreshToken string) (*v1Auth.LogoutResponse, error)
	ResetPassword(userId, oldPassword, newPassword string) (*v1Auth.ResetPasswordResponse, error)
	ForgotPassword(email string) (*v1Auth.ForgotPasswordResponse, error)
	ConfirmForgotPassword(token, userId string) (*v1Auth.ConfirmForgotPasswordResponse, error)
	SaveRouteResource(resources []*v1Auth.ResourceItem) (*v1Auth.SaveRouteResourceResponse, error)
}

type AuthClient struct {
	logger     *loggers.LoggerZap
	config     *settings.AuthService
	authClient v1Auth.AuthServiceClient
}

func NewAuthClient() IAuthClient {
	logger := global.Logger
	config := global.Config.AuthService

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorString("Failed to connect to AuthService", zap.String("host", config.Host), zap.Int("port", config.Port), zap.Error(err))
		return nil
	}

	client := v1Auth.NewAuthServiceClient(conn)
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

func (a *AuthClient) Login(username, password string) (*v1Auth.LoginResponse, error) {
	req := &v1Auth.LoginRequest{
		Username: username,
		Password: password,
	}
	resp, err := a.authClient.Login(context.Background(), req)
	if err != nil {
		a.logger.ErrorString("Login failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) Register(email, password string) (*v1Auth.RegisterResponse, error) {
	req := &v1Auth.RegisterRequest{
		Email:    email,
		Password: password,
	}
	resp, err := a.authClient.Register(context.Background(), req)
	if err != nil {
		a.logger.ErrorString("Register failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) ConfirmEmail(email, token, password string) (*v1Auth.ConfirmEmailResponse, error) {
	req := &v1Auth.ConfirmEmailRequest{
		Email:    email,
		Token:    token,
		Password: password,
	}
	resp, err := a.authClient.ConfirmEmail(context.Background(), req)
	if err != nil {
		a.logger.ErrorString("ConfirmEmail failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) Logout(accessToken, refreshToken string) (*v1Auth.LogoutResponse, error) {
	req := &v1Auth.LogoutRequest{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	resp, err := a.authClient.Logout(context.Background(), req)
	if err != nil {
		a.logger.ErrorString("Logout failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) ResetPassword(userId, oldPassword, newPassword string) (*v1Auth.ResetPasswordResponse, error) {
	req := &v1Auth.ResetPasswordRequest{
		UserId:      userId,
		OldPassword: oldPassword,
		NewPassword: newPassword,
	}
	resp, err := a.authClient.ResetPassword(context.Background(), req)
	if err != nil {
		a.logger.ErrorString("ResetPassword failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) ForgotPassword(email string) (*v1Auth.ForgotPasswordResponse, error) {
	req := &v1Auth.ForgotPasswordRequest{
		Email: email,
	}
	resp, err := a.authClient.ForgotPassword(context.Background(), req)
	if err != nil {
		a.logger.ErrorString("ForgotPassword failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) ConfirmForgotPassword(token, userId string) (*v1Auth.ConfirmForgotPasswordResponse, error) {
	req := &v1Auth.ConfirmForgotPasswordRequest{
		UserId: userId,
		Token:  token,
	}
	resp, err := a.authClient.ConfirmForgotPassword(context.Background(), req)
	if err != nil {
		a.logger.ErrorString("ConfirmForgotPassword failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (a *AuthClient) SaveRouteResource(resources []*v1Auth.ResourceItem) (*v1Auth.SaveRouteResourceResponse, error) {
	req := &v1Auth.SaveRouteResourceRequest{
		Items: resources,
	}
	resp, err := a.authClient.SaveRouteResource(context.Background(), req)
	if err != nil {
		a.logger.ErrorString("SaveRouteResource failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
