package auth_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type tokenClient struct {
	logger      log.Logger
	tokenClient auth.TokenServiceClient
}

func (t *tokenClient) RefreshToken(c *gin.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := t.tokenClient.RefreshToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (t *tokenClient) RevokeToken(c *gin.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := t.tokenClient.RevokeToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
