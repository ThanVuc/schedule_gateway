package user_client

import (
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/pkg/settings"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/user"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	UserClient interface {
		GetUserProfile(c *gin.Context, req *user.GetUserProfileRequest) (*user.GetUserProfileResponse, error)
		UpdateUserProfile(c *gin.Context, req *user.UpdateUserProfileRequest) (*common.EmptyResponse, error)
	}
)

func getConn(baseConfig settings.GrpcBase) *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", baseConfig.GetHost(), baseConfig.GetPort()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to gRPC server: " + err.Error())
	}
	return conn
}

func NewUserClient() UserClient {
	conn := getConn(&global.Config.UserService)

	client := user.NewUserServiceClient(conn)
	if client == nil {
		panic("Failed to create UserService client at " + fmt.Sprintf("%s:%d", global.Config.AuthService.GetHost(), global.Config.AuthService.GetPort()))
	}

	return &userClient{
		userClient: client,
	}
}
