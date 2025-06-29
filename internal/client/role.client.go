package client

import (
	"context"
	"schedule_gateway/internal/grpc/auth"
	"schedule_gateway/pkg/loggers"

	"go.uber.org/zap"
)

type roleClient struct {
	logger     *loggers.LoggerZap
	roleClient auth.RoleServiceClient
}

func (r *roleClient) GetRoles(ctx context.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error) {
	resp, err := r.roleClient.GetRoles(ctx, req)
	if err != nil {
		r.logger.ErrorString("GetRoles failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) CreateRole(ctx context.Context, req *auth.CreateRoleRequest) (*auth.CreateRoleResponse, error) {
	resp, err := r.roleClient.CreateRole(ctx, req)
	if err != nil {
		r.logger.ErrorString("CreateRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) UpdateRole(ctx context.Context, req *auth.UpdateRoleRequest) (*auth.UpdateRoleResponse, error) {
	resp, err := r.roleClient.UpdateRole(ctx, req)
	if err != nil {
		r.logger.ErrorString("UpdateRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error) {
	resp, err := r.roleClient.DeleteRole(ctx, req)
	if err != nil {
		r.logger.ErrorString("DeleteRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error) {
	resp, err := r.roleClient.DisableOrEnableRole(ctx, req)
	if err != nil {
		r.logger.ErrorString("DisableOrEnableRole failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) AssignRoleToUser(ctx context.Context, req *auth.AssignRoleRequest) (*auth.AssignRoleResponse, error) {
	resp, err := r.roleClient.AssignRoleToUser(ctx, req)
	if err != nil {
		r.logger.ErrorString("AssignRoleToUser failed", zap.Error(err))
		return nil, err
	}
	return resp, nil
}
