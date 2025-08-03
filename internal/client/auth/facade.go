package auth_client

import (
	"context"
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/pkg/settings"
	"schedule_gateway/proto/auth"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/user"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	AuthClient interface {
		LoginWithGoogle(c *gin.Context, req *auth.LoginWithGoogleRequest) (*auth.LoginWithGoogleResponse, error)
		Logout(c *gin.Context, req *auth.LogoutRequest) (*common.EmptyResponse, error)
		SaveRouteResource(context context.Context, req *auth.SaveRouteResourceRequest) (*auth.SaveRouteResourceResponse, error)
	}

	PermissionClient interface {
		GetPermissions(c *gin.Context, req *auth.GetPermissionsRequest) (*auth.GetPermissionsResponse, error)
		UpsertPermission(c *gin.Context, req *auth.UpsertPermissionRequest) (*auth.UpsertPermissionResponse, error)
		DeletePermission(c *gin.Context, req *auth.DeletePermissionRequest) (*auth.DeletePermissionResponse, error)
		GetResources(c *gin.Context, req *auth.GetResourcesRequest) (*auth.GetResourcesResponse, error)
		GetActions(c *gin.Context, req *auth.GetActionsRequest) (*auth.GetActionsResponse, error)
		GetPermission(c *gin.Context, req *auth.GetPermissionRequest) (*auth.GetPermissionResponse, error)
	}

	RoleClient interface {
		GetRoles(c *gin.Context, req *auth.GetRolesRequest) (*auth.GetRolesResponse, error)
		GetRole(c *gin.Context, req *auth.GetRoleRequest) (*auth.GetRoleResponse, error)
		DeleteRole(c *gin.Context, req *auth.DeleteRoleRequest) (*auth.DeleteRoleResponse, error)
		DisableOrEnableRole(c *gin.Context, req *auth.DisableOrEnableRoleRequest) (*auth.DisableOrEnableRoleResponse, error)
		UpsertRole(c *gin.Context, req *auth.UpsertRoleRequest) (*auth.UpsertRoleResponse, error)
	}

	TokenClient interface {
		RefreshToken(c *gin.Context, req *auth.RefreshTokenRequest) (*auth.RefreshTokenResponse, error)
		RevokeToken(c *gin.Context, req *auth.RevokeTokenRequest) (*auth.RevokeTokenResponse, error)
	}

	UserClient interface {
	}
)

func getConn(baseConfig settings.GrpcBase) *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", baseConfig.GetHost(), baseConfig.GetPort()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to gRPC server: " + err.Error())
	}
	return conn
}

func NewAuthClient() AuthClient {
	conn := getConn(&global.Config.AuthService)

	client := auth.NewAuthServiceClient(conn)
	if client == nil {
		panic("Failed to create AuthService client at " + fmt.Sprintf("%s:%d", global.Config.AuthService.GetHost(), global.Config.AuthService.GetPort()))
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
		panic("Failed to create PermissionService client at " + fmt.Sprintf("%s:%d", global.Config.AuthService.GetHost(), global.Config.AuthService.GetPort()))
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
		panic("Failed to create RoleService client at " + fmt.Sprintf("%s:%d", global.Config.AuthService.GetHost(), global.Config.AuthService.GetPort()))
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
		panic("Failed to create TokenService client at " + fmt.Sprintf("%s:%d", global.Config.AuthService.GetHost(), global.Config.AuthService.GetPort()))
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
		panic("Failed to create UserService client at " + fmt.Sprintf("%s:%d", global.Config.UserService.GetHost(), global.Config.UserService.GetPort()))
	}

	return &userClient{
		logger:     global.Logger,
		userClient: client,
	}
}
