package client

import (
	"context"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/proto/auth"
)

type authClient struct {
	logger     *loggers.LoggerZap
	authClient auth.AuthServiceClient
}

func (a *authClient) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	resp, err := a.authClient.Login(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authClient) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	resp, err := a.authClient.Register(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authClient) ConfirmEmail(ctx context.Context, req *auth.ConfirmEmailRequest) (*auth.ConfirmEmailResponse, error) {
	resp, err := a.authClient.ConfirmEmail(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authClient) Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error) {
	resp, err := a.authClient.Logout(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authClient) ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error) {
	resp, err := a.authClient.ResetPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authClient) ForgotPassword(ctx context.Context, req *auth.ForgotPasswordRequest) (*auth.ForgotPasswordResponse, error) {
	resp, err := a.authClient.ForgotPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authClient) ConfirmForgotPassword(ctx context.Context, req *auth.ConfirmForgotPasswordRequest) (*auth.ConfirmForgotPasswordResponse, error) {
	resp, err := a.authClient.ConfirmForgotPassword(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authClient) SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error) {
	resp, err := a.authClient.SaveRouteResource(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
