package client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/proto/auth"
	"schedule_gateway/proto/common"

	"github.com/gin-gonic/gin"
)

type authClient struct {
	logger     *loggers.LoggerZap
	authClient auth.AuthServiceClient
}

func (a *authClient) LoginWithGoogle(c *gin.Context, req *auth.LoginWithGoogleRequest) (*auth.LoginWithGoogleResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := a.authClient.LoginWithGoogle(ctx, req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authClient) Logout(c *gin.Context, req *auth.LogoutRequest) (*common.EmptyResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := a.authClient.Logout(ctx, req)
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
