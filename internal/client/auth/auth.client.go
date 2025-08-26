package auth_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/auth"
	"schedule_gateway/proto/common"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type authClient struct {
	logger     log.Logger
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

func (a *authClient) RefreshToken(c *gin.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := a.authClient.RefreshToken(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (a *authClient) CheckPermission(c *gin.Context, req *auth.CheckPermissionRequest) (*auth.CheckPermissionResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := a.authClient.CheckPermission(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
