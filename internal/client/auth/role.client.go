package auth_service

import (
	"context"
	"fmt"
	"schedule_gateway/global"
	v1Role "schedule_gateway/internal/grpc/role.v1"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/pkg/settings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type IRoleClient interface {
	GetRoles(userID string) (*v1Role.GetRolesResponse, error)
	CreateRole(name, description string) (*v1Role.CreateRoleResponse, error)
	UpdateRole(roleID, name, description string) (*v1Role.UpdateRoleResponse, error)
	DeleteRole(roleID string) (*v1Role.DeleteRoleResponse, error)
	DisableOrEnableRole(roleID string, disable bool) (*v1Role.DisableOrEnableRoleResponse, error)
	AssignRoleToUser(userID, roleID string) (*v1Role.AssignRoleResponse, error)
}

type RoleClient struct {
	logger     *loggers.LoggerZap
	config     *settings.AuthService
	roleClient v1Role.RoleServiceClient
}

func NewRoleClient() IRoleClient {
	logger := global.Logger
	config := global.Config.AuthService

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.ErrorString("Failed to connect to RoleService", zap.String("host", config.Host), zap.Int("port", config.Port), zap.Error(err))
		return nil
	}

	client := v1Role.NewRoleServiceClient(conn)
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

func (r *RoleClient) GetRoles(userID string) (*v1Role.GetRolesResponse, error) {
	req := &v1Role.GetRolesRequest{
		UserId: userID,
	}
	resp, err := r.roleClient.GetRoles(context.Background(), req)
	if err != nil {
		r.logger.ErrorString("GetRoles failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *RoleClient) CreateRole(name, description string) (*v1Role.CreateRoleResponse, error) {
	req := &v1Role.CreateRoleRequest{
		Name:        name,
		Description: description,
	}
	resp, err := r.roleClient.CreateRole(context.Background(), req)
	if err != nil {
		r.logger.ErrorString("CreateRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *RoleClient) UpdateRole(roleID, name, description string) (*v1Role.UpdateRoleResponse, error) {
	req := &v1Role.UpdateRoleRequest{
		RoleId:      roleID,
		Name:        name,
		Description: description,
	}
	resp, err := r.roleClient.UpdateRole(context.Background(), req)
	if err != nil {
		r.logger.ErrorString("UpdateRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *RoleClient) DeleteRole(roleID string) (*v1Role.DeleteRoleResponse, error) {
	req := &v1Role.DeleteRoleRequest{
		RoleId: roleID,
	}
	resp, err := r.roleClient.DeleteRole(context.Background(), req)
	if err != nil {
		r.logger.ErrorString("DeleteRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *RoleClient) DisableOrEnableRole(roleID string, disable bool) (*v1Role.DisableOrEnableRoleResponse, error) {
	req := &v1Role.DisableOrEnableRoleRequest{
		RoleId:  roleID,
		Disable: disable,
	}
	resp, err := r.roleClient.DisableOrEnableRole(context.Background(), req)
	if err != nil {
		r.logger.ErrorString("DisableOrEnableRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *RoleClient) AssignRoleToUser(userID, roleID string) (*v1Role.AssignRoleResponse, error) {
	req := &v1Role.AssignRoleRequest{
		UserId: userID,
		RoleId: roleID,
	}
	resp, err := r.roleClient.AssignRoleToUser(context.Background(), req)
	if err != nil {
		r.logger.ErrorString("AssignRoleToUser failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
