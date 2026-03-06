package team_client

import (
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/pkg/settings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	GroupClient  interface{}
	SprintClient interface{}
	WorkClient   interface{}
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
	client := (conn)
	if client == nil {
		panic("Failed to create TeamService client at " + fmt.Sprintf("%s:%d", global.Config.TeamService.GetHost(), global.Config.TeamService.GetPort()))
	}
	return &groupClient{
		teamClient: client,
	}
}

func NewSprintClient() SprintClient {
	conn := getConn(&global.Config.TeamService)
	client := (conn)

	if client == nil {
		panic("Failed to create TeamService client at " + fmt.Sprintf("%s:%d", global.Config.TeamService.GetHost(), global.Config.TeamService.GetPort()))
	}
	return &sprintClient{
		teamClient: client,
	}
}

func NewWorkClient() WorkClient {
	conn := getConn(&global.Config.TeamService)
	client := (conn)

	if client == nil {
		panic("Failed to create TeamService client at " + fmt.Sprintf("%s:%d", global.Config.TeamService.GetHost(), global.Config.TeamService.GetPort()))
	}

	return &workClient{
		teamClient: client,
	}
}
