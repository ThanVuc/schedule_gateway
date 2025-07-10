package client

import (
	"context"
	"schedule_gateway/pkg/loggers"
	"schedule_gateway/proto/auth"
)

type roleClient struct {
	logger     *loggers.LoggerZap
	roleClient auth.RoleServiceClient
}

func (r *roleClient) GetRoles(ctx context.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error) {
	resp, err := r.roleClient.GetRoles(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) GetRole(ctx context.Context, req *auth.GetRoleRequest) (*auth.GetRoleResponse, error) {
	resp, err := r.roleClient.GetRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error) {
	resp, err := r.roleClient.DeleteRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error) {
	resp, err := r.roleClient.DisableOrEnableRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (r *roleClient) UpsertRole(ctx context.Context, req *auth.UpsertRoleRequest) (*auth.UpsertRoleResponse, error) {
	resp, err := r.roleClient.UpsertRole(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
