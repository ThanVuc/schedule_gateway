package team_client

import (
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/pkg/settings"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/team_service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GroupClient interface {
		Ping(c *gin.Context, req *common.EmptyRequest) (*common.EmptyResponse, error)
		CreateGroup(c *gin.Context, req *team_service.CreateGroupRequest) (*team_service.CreateGroupResponse, error)
		GetGroup(c *gin.Context, req *common.IDRequest) (*team_service.GetGroupResponse, error)
	}
	SprintClient interface {
		CreateSprint(c *gin.Context, req *team_service.CreateSprintRequest) (*team_service.CreateSprintResponse, error)
		GetSprint(c *gin.Context, req *common.IDRequest) (*team_service.GetSprintResponse, error)
		ListSprints(c *gin.Context, req *team_service.ListSprintsRequest) (*team_service.ListSprintsResponse, error)
		UpdateSprint(c *gin.Context, req *team_service.UpdateSprintRequest) (*team_service.UpdateSprintResponse, error)
		UpdateSprintStatus(c *gin.Context, req *team_service.UpdateSprintStatusRequest) (*team_service.UpdateSprintStatusResponse, error)
		DeleteSprint(c *gin.Context, req *common.IDRequest) (*team_service.DeleteSprintResponse, error)
	}
	WorkClient interface{}
)

func getConn(baseConfig settings.GrpcBase) *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", baseConfig.GetHost(), baseConfig.GetPort()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to gRPC server: " + err.Error())
	}
	return conn
}

func NewGroupClient() GroupClient {
	conn := getConn(&global.Config.TeamService)
	client := team_service.NewGroupServiceClient(conn)
	if client == nil {
		panic("Failed to create TeamService client at " + fmt.Sprintf("%s:%d", global.Config.TeamService.GetHost(), global.Config.TeamService.GetPort()))
	}
	return &groupClient{
		groupClient: client,
	}
}

func NewSprintClient() SprintClient {
	conn := getConn(&global.Config.TeamService)
	client := team_service.NewSprintServiceClient(conn)

	if client == nil {
		panic("Failed to create TeamService client at " + fmt.Sprintf("%s:%d", global.Config.TeamService.GetHost(), global.Config.TeamService.GetPort()))
	}
	return &sprintClient{
		sprintClient: client,
	}
}

func NewWorkClient() WorkClient {
	conn := getConn(&global.Config.TeamService)
	client := team_service.NewWorkServiceClient(conn)

	if client == nil {
		panic("Failed to create TeamService client at " + fmt.Sprintf("%s:%d", global.Config.TeamService.GetHost(), global.Config.TeamService.GetPort()))
	}

	return &workClient{
		workClient: client,
	}
}
