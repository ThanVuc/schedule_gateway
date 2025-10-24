package auth_client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/auth"
	"schedule_gateway/proto/common"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type userClient struct {
	logger     log.Logger
	userClient auth.UserServiceClient
}

func (uc *userClient) AssignRoleToUser(c *gin.Context, req *auth.AssignRoleToUserRequest) (*common.EmptyResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := uc.userClient.AssignRoleToUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *userClient) GetUsers(c *gin.Context, req *auth.GetUsersRequest) (*auth.GetUsersResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := p.userClient.GetUsers(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *userClient) GetUser(c *gin.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := p.userClient.GetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *userClient) LockOrUnLockUser(c *gin.Context, req *auth.LockUserRequest) (*common.EmptyResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := p.userClient.LockOrUnLockUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *userClient) PresignUrlForAvatarUpsert(c *gin.Context, req *auth.PresignUrlRequest) (*auth.PresignRequestUrlForAvatarUpsertResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := p.userClient.PresignUrlForAvatarUpsert(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
