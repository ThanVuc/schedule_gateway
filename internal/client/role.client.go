package client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type roleClient struct {
	logger     log.Logger
	roleClient auth.RoleServiceClient
}

func (r *roleClient) GetRoles(c *gin.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := r.roleClient.GetRoles(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) GetRole(c *gin.Context, req *auth.GetRoleRequest) (*auth.GetRoleResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := r.roleClient.GetRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) DeleteRole(c *gin.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := r.roleClient.DeleteRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) DisableOrEnableRole(c *gin.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := r.roleClient.DisableOrEnableRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) UpsertRole(c *gin.Context, req *auth.UpsertRoleRequest) (*auth.UpsertRoleResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := r.roleClient.UpsertRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
