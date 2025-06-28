package client

import (
	"context"
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/internal/grpc/auth"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/settings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IPermissionClient interface {
	GetPermissions(ctx context.Context, req *auth.GetPermissionsRequest) (*auth.GetPermissionsResponse, error)
	CreatePermission(ctx context.Context, req *auth.CreatePermissionRequest) (*auth.CreatePermissionResponse, error)
	UpdatePermission(ctx context.Context, req *auth.UpdatePermissionRequest) (*auth.UpdatePermissionResponse, error)
	DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error)
	AssignPermissionToRole(ctx context.Context, req *auth.AssignPermissionRequest) (*auth.AssignPermissionResponse, error)
}

type PermissionClient struct {
	logger           *loggers.LoggerZap
	config           *settings.AuthService
	permissionClient auth.PermissionServiceClient
}

func NewPermissionClient() IPermissionClient {
	logger := global.Logger
	config := global.Config.AuthService

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorString("Failed to connect to PermissionService", zap.String("host", config.Host), zap.Int("port", config.Port), zap.Error(err))
		return nil
	}

	client := auth.NewPermissionServiceClient(conn)
	if client == nil {
		logger.ErrorString("Failed to create PermissionService client", zap.String("host", config.Host), zap.Int("port", config.Port))
		return nil
	}

	return &PermissionClient{
		logger:           logger,
		config:           &config,
		permissionClient: client,
	}
}

func (p *PermissionClient) GetPermissions(ctx context.Context, req *auth.GetPermissionsRequest) (*auth.GetPermissionsResponse, error) {
	resp, err := p.permissionClient.GetPermissions(ctx, req)
	if err != nil {
		p.logger.ErrorString("GetPermissions failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *PermissionClient) CreatePermission(ctx context.Context, req *auth.CreatePermissionRequest) (*auth.CreatePermissionResponse, error) {
	resp, err := p.permissionClient.CreatePermission(ctx, req)
	if err != nil {
		p.logger.ErrorString("CreatePermission failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *PermissionClient) UpdatePermission(ctx context.Context, req *auth.UpdatePermissionRequest) (*auth.UpdatePermissionResponse, error) {
	resp, err := p.permissionClient.UpdatePermission(ctx, req)
	if err != nil {
		p.logger.ErrorString("UpdatePermission failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *PermissionClient) DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error) {
	resp, err := p.permissionClient.DeletePermission(ctx, req)
	if err != nil {
		p.logger.ErrorString("DeletePermission failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *PermissionClient) AssignPermissionToRole(ctx context.Context, req *auth.AssignPermissionRequest) (*auth.AssignPermissionResponse, error) {
	resp, err := p.permissionClient.AssignPermissionToRole(ctx, req)
	if err != nil {
		p.logger.ErrorString("AssignPermissionToRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
