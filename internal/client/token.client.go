package client

import (
	"context"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/proto/auth"
)

type tokenClient struct {
	logger      *loggers.LoggerZap
	tokenClient auth.TokenServiceClient
}

func (t *tokenClient) RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	resp, err := t.tokenClient.RefreshToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *tokenClient) RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error) {
	resp, err := t.tokenClient.RevokeToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
