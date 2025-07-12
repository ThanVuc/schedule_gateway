package client

import (
	"context"
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/pkg/response"
	"schedule_gateway/pkg/settings"
	"schedule_gateway/proto/auth"
	"schedule_gateway/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	AuthClient interface {
		Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error)
		Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error)
		ConfirmEmail(ctx context.Context, req *auth.ConfirmEmailRequest) (*auth.ConfirmEmailResponse, error)
		Logout(ctx context.Context, req *auth.LogoutRequest) (*auth.LogoutResponse, error)
		ResetPassword(ctx context.Context, req *auth.ResetPasswordRequest) (*auth.ResetPasswordResponse, error)
		ForgotPassword(ctx context.Context, req *auth.ForgotPasswordRequest) (*auth.ForgotPasswordResponse, error)
		ConfirmForgotPassword(ctx context.Context, req *auth.ConfirmForgotPasswordRequest) (*auth.ConfirmForgotPasswordResponse, error)
		SaveRouteResource(ctx context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error)
	}

	PermissionClient interface {
		GetPermissions(ctx context.Context, req *auth.GetPermissionsRequest) (*auth.GetPermissionsResponse, error)
		UpsertPermission(ctx context.Context, req *auth.UpsertPermissionRequest) (*auth.UpsertPermissionResponse, error)
		DeletePermission(ctx context.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error)
		GetResources(ctx context.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error)
		GetActions(ctx context.Context, req *auth.GetActionsRequest) (*auth.GetActionsResponse, error)
		GetPermission(ctx context.Context, req *auth.GetPermissionRequest) (*auth.GetPermissionResponse, error)
	}

	RoleClient interface {
		GetRoles(ctx context.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error)
		GetRole(ctx context.Context, req *auth.GetRoleRequest) (*auth.GetRoleResponse, error)
		DeleteRole(ctx context.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error)
		DisableOrEnableRole(ctx context.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error)
		UpsertRole(ctx context.Context, req *auth.UpsertRoleRequest) (*auth.UpsertRoleResponse, error)
	}

	TokenClient interface {
		RefreshToken(ctx context.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error)
		RevokeToken(ctx context.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error)
	}

	UserClient interface {
		GetUserInfo(ctx context.Context, req *user.GetUserInfoRequest) (*user.GetUserInfoResponse, error)
		UpdateUserInfo(ctx context.Context, req *user.UpdateUserInfoRequest) (*user.UpdateUserInfoResponse, error)
	}
)

func getConn(baseConfig settings.GrpcBase) *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", baseConfig.GetHost(), baseConfig.GetPort()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(response.InternalServerError(fmt.Sprintf("Failed to connect to service at %s:%d", baseConfig.GetHost(), baseConfig.GetPort())))
	}
	return conn
}

func NewAuthClient() AuthClient {
	conn := getConn(&global.Config.AuthService)

	client := auth.NewAuthServiceClient(conn)
	if client == nil {
		panic(response.InternalServerError(fmt.Sprintf("Failed to create AuthService client at %s:%d", global.Config.AuthService.GetHost(), global.Config.AuthService.GetPort())))
	}

	return &authClient{
		logger:     global.Logger,
		authClient: client,
	}
}

func NewPermissionClient() PermissionClient {
	conn := getConn(&global.Config.AuthService)

	client := auth.NewPermissionServiceClient(conn)
	if client == nil {
		panic(response.InternalServerError(fmt.Sprintf("Failed to create PermissionService client at %s:%d", global.Config.AuthService.GetHost(), global.Config.AuthService.GetPort())))
	}

	return &permissionClient{
		logger:           global.Logger,
		permissionClient: client,
	}
}

func NewRoleClient() RoleClient {
	conn := getConn(&global.Config.AuthService)

	client := auth.NewRoleServiceClient(conn)
	if client == nil {
		panic(response.InternalServerError(fmt.Sprintf("Failed to create RoleService client at %s:%d", global.Config.AuthService.GetHost(), global.Config.AuthService.GetPort())))
	}

	return &roleClient{
		logger:     global.Logger,
		roleClient: client,
	}
}

func NewTokenClient() TokenClient {
	conn := getConn(&global.Config.AuthService)

	client := auth.NewTokenServiceClient(conn)
	if client == nil {
		panic(response.InternalServerError(fmt.Sprintf("Failed to create TokenService client at %s:%d", global.Config.AuthService.GetHost(), global.Config.AuthService.GetPort())))
	}

	return &tokenClient{
		logger:      global.Logger,
		tokenClient: client,
	}
}

func NewUserClient() UserClient {
	conn := getConn(&global.Config.UserService)

	client := user.NewUserServiceClient(conn)
	if client == nil {
		panic(response.InternalServerError(fmt.Sprintf("Failed to create UserService client at %s:%d", global.Config.UserService.GetHost(), global.Config.UserService.GetPort())))
	}

	return &userClient{
		logger:     global.Logger,
		userClient: client,
	}
}
