package client

import (
	"context"
	"schedule_gateway/internal/utils"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
	"github.com/thanvuc/go-core-lib/log"
)

type permissionClient struct {
	logger           log.Logger
	permissionClient auth.PermissionServiceClient
}

func (p *permissionClient) GetPermissions(c *gin.Context, req *auth.GetPermissionsRequest) (*auth.GetPermissionsResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := p.permissionClient.GetPermissions(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *permissionClient) UpsertPermission(c *gin.Context, req *auth.UpsertPermissionRequest) (*auth.UpsertPermissionResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := p.permissionClient.UpsertPermission(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *permissionClient) DeletePermission(c *gin.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := p.permissionClient.DeletePermission(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *permissionClient) GetResources(c *gin.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := p.permissionClient.GetResources(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *permissionClient) GetActions(c *gin.Context, req *auth.GetActionsRequest) (*auth.GetActionsResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := p.permissionClient.GetActions(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *permissionClient) GetPermission(c *gin.Context, req *auth.GetPermissionRequest) (*auth.GetPermissionResponse, error) {
	ctx := context.Background()
	ctx = utils.WithRequestID(ctx, c.GetString("request-id"))

	resp, err := p.permissionClient.GetPermission(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
