package client

import (
	"context"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/proto/auth"

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

func (p *permissionClient) UpsertPermission(ctx context.Context, req *auth.UpsertPermissionRequest) (*auth.UpsertPermissionResponse, error) {
	resp, err := p.permissionClient.UpsertPermission(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *permissionClient) DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error) {
	resp, err := p.permissionClient.DeletePermission(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *permissionClient) AssignPermissionToRole(ctx context.Context, req *auth.AssignPermissionRequest) (*auth.AssignPermissionResponse, error) {
	resp, err := p.permissionClient.AssignPermissionToRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (p *permissionClient) GetResources(ctx context.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error) {
	resp, err := p.permissionClient.GetResources(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (p *permissionClient) GetActions(ctx context.Context, req *auth.GetActionsRequest) (*auth.GetActionsResponse, error) {
	resp, err := p.permissionClient.GetActions(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
