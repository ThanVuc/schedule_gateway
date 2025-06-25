package client

import (
	"context"
	"fmt"
	"schedule_gateway/global"
	v1Permission "schedule_gateway/internal/grpc/permission.v1"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/settings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IPermissionClient interface {
	GetPermissions(userID string) (*v1Permission.GetPermissionsResponse, error)
	CreatePermission(name, description string) (*v1Permission.CreatePermissionResponse, error)
	UpdatePermission(permissionID, name, description string) (*v1Permission.UpdatePermissionResponse, error)
	DeletePermission(permissionID string) (*v1Permission.DeletePermissionResponse, error)
	AssignPermissionToRole(permissionID, roleID string) (*v1Permission.AssignPermissionResponse, error)
}

type PermissionClient struct {
	logger           *loggers.LoggerZap
	config           *settings.AuthService
	permissionClient v1Permission.PermissionServiceClient
}

func NewPermissionClient() IPermissionClient {
	logger := global.Logger
	config := global.Config.AuthService

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorString("Failed to connect to PermissionService", zap.String("host", config.Host), zap.Int("port", config.Port), zap.Error(err))
		return nil
	}

	client := v1Permission.NewPermissionServiceClient(conn)
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

func (p *PermissionClient) GetPermissions(userID string) (*v1Permission.GetPermissionsResponse, error) {
	req := &v1Permission.GetPermissionsRequest{
		UserId: userID,
	}
	resp, err := p.permissionClient.GetPermissions(context.Background(), req)
	if err != nil {
		p.logger.ErrorString("GetPermissions failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *PermissionClient) CreatePermission(name, description string) (*v1Permission.CreatePermissionResponse, error) {
	req := &v1Permission.CreatePermissionRequest{
		Name:        name,
		Description: description,
	}
	resp, err := p.permissionClient.CreatePermission(context.Background(), req)
	if err != nil {
		p.logger.ErrorString("CreatePermission failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *PermissionClient) UpdatePermission(permissionID, name, description string) (*v1Permission.UpdatePermissionResponse, error) {
	req := &v1Permission.UpdatePermissionRequest{
		PermissionId: permissionID,
		Name:         name,
		Description:  description,
	}
	resp, err := p.permissionClient.UpdatePermission(context.Background(), req)
	if err != nil {
		p.logger.ErrorString("UpdatePermission failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *PermissionClient) DeletePermission(permissionID string) (*v1Permission.DeletePermissionResponse, error) {
	req := &v1Permission.DeletePermissionRequest{
		PermissionId: permissionID,
	}
	resp, err := p.permissionClient.DeletePermission(context.Background(), req)
	if err != nil {
		p.logger.ErrorString("DeletePermission failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (p *PermissionClient) AssignPermissionToRole(permissionID, roleID string) (*v1Permission.AssignPermissionResponse, error) {
	req := &v1Permission.AssignPermissionRequest{
		PermissionId: permissionID,
		RoleId:       roleID,
	}
	resp, err := p.permissionClient.AssignPermissionToRole(context.Background(), req)
	if err != nil {
		p.logger.ErrorString("AssignPermissionToRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
