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

type IRoleClient interface {
	GetRoles(ctx context.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error)
	CreateRole(ctx context.Context, req *auth.CreateRoleRequest) (*auth.CreateRoleResponse, error)
	UpdateRole(ctx context.Context, req *auth.UpdateRoleRequest) (*auth.UpdateRoleResponse, error)
	DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error)
	DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error)
	AssignRoleToUser(ctx context.Context, req *auth.AssignRoleRequest) (*auth.AssignRoleResponse, error)
}

type RoleClient struct {
	logger     *loggers.LoggerZap
	config     *settings.AuthService
	roleClient auth.RoleServiceClient
}

func NewRoleClient() IRoleClient {
	logger := global.Logger
	config := global.Config.AuthService

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorString("Failed to connect to RoleService", zap.String("host", config.Host), zap.Int("port", config.Port), zap.Error(err))
		return nil
	}

	client := auth.NewRoleServiceClient(conn)
	if client == nil {
		logger.ErrorString("Failed to create RoleService client", zap.String("host", config.Host), zap.Int("port", config.Port))
		return nil
	}

	return &RoleClient{
		logger:     logger,
		config:     &config,
		roleClient: client,
	}
}

func (r *RoleClient) GetRoles(ctx context.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error) {
	resp, err := r.roleClient.GetRoles(ctx, req)
	if err != nil {
		r.logger.ErrorString("GetRoles failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *RoleClient) CreateRole(ctx context.Context, req *auth.CreateRoleRequest) (*auth.CreateRoleResponse, error) {
	resp, err := r.roleClient.CreateRole(ctx, req)
	if err != nil {
		r.logger.ErrorString("CreateRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *RoleClient) UpdateRole(ctx context.Context, req *auth.UpdateRoleRequest) (*auth.UpdateRoleResponse, error) {
	resp, err := r.roleClient.UpdateRole(ctx, req)
	if err != nil {
		r.logger.ErrorString("UpdateRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *RoleClient) DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error) {
	resp, err := r.roleClient.DeleteRole(ctx, req)
	if err != nil {
		r.logger.ErrorString("DeleteRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *RoleClient) DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error) {
	resp, err := r.roleClient.DisableOrEnableRole(ctx, req)
	if err != nil {
		r.logger.ErrorString("DisableOrEnableRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *RoleClient) AssignRoleToUser(ctx context.Context, req *auth.AssignRoleRequest) (*auth.AssignRoleResponse, error) {
	resp, err := r.roleClient.AssignRoleToUser(ctx, req)
	if err != nil {
		r.logger.ErrorString("AssignRoleToUser failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
