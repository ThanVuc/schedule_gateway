package client

import (
	"context"
	"schedule_gateway/internal/grpc/auth"
	"schedule_gateway/pkg/loggers"

	"go.uber.org/zap"
)

type permissionClient struct {
	logger           *loggers.LoggerZap
	permissionClient auth.PermissionServiceClient
}

func (p *permissionClient) GetPermissions(ctx context.Context, req *auth.GetPermissionsRequest) (*auth.GetPermissionsResponse, error) {
	resp, err := p.permissionClient.GetPermissions(ctx, req)
	if err != nil {
		p.logger.ErrorString("GetPermissions failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *permissionClient) CreatePermission(ctx context.Context, req *auth.CreatePermissionRequest) (*auth.CreatePermissionResponse, error) {
	resp, err := p.permissionClient.CreatePermission(ctx, req)
	if err != nil {
		p.logger.ErrorString("CreatePermission failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *permissionClient) UpdatePermission(ctx context.Context, req *auth.UpdatePermissionRequest) (*auth.UpdatePermissionResponse, error) {
	resp, err := p.permissionClient.UpdatePermission(ctx, req)
	if err != nil {
		p.logger.ErrorString("UpdatePermission failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *permissionClient) DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error) {
	resp, err := p.permissionClient.DeletePermission(ctx, req)
	if err != nil {
		p.logger.ErrorString("DeletePermission failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *permissionClient) AssignPermissionToRole(ctx context.Context, req *auth.AssignPermissionRequest) (*auth.AssignPermissionResponse, error) {
	resp, err := p.permissionClient.AssignPermissionToRole(ctx, req)
	if err != nil {
		p.logger.ErrorString("AssignPermissionToRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *permissionClient) GetResources(ctx context.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error) {
	resp, err := p.permissionClient.GetResources(ctx, req)
	if err != nil {
		p.logger.ErrorString("GetResources failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
