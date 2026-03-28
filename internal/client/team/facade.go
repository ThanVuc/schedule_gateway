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
		ListGroups(c *gin.Context, req *common.IDRequest) (*team_service.ListGroupsResponse, error)
		GetSimpleUserByGroupID(c *gin.Context, req *common.IDRequest) (*team_service.GetSimpleUserByGroupIDResponse, error)
		UpdateGroup(c *gin.Context, req *team_service.UpdateGroupRequest) (*team_service.UpdateGroupResponse, error)
		DeleteGroup(c *gin.Context, req *common.IDRequest) (*team_service.DeleteGroupResponse, error)
		ListMembers(c *gin.Context, req *team_service.ListMembersRequest) (*team_service.ListMembersResponse, error)
		UpdateMemberRole(c *gin.Context, req *team_service.UpdateMemberRoleRequest) (*team_service.UpdateMemberRoleResponse, error)
		RemoveMember(c *gin.Context, req *team_service.RemoveMemberRequest) (*team_service.RemoveMemberResponse, error)
		CreateInvite(c *gin.Context, req *team_service.CreateInviteRequest) (*team_service.CreateInviteResponse, error)
	}
	SprintClient interface {
		CreateSprint(c *gin.Context, req *team_service.CreateSprintRequest) (*team_service.CreateSprintResponse, error)
		GetSprint(c *gin.Context, req *common.IDRequest) (*team_service.GetSprintResponse, error)
		GetSimpleSprints(c *gin.Context, req *common.IDRequest) (*team_service.GetSimpleSprintsResponse, error)
		ExportSprint(c *gin.Context, req *common.IDRequest) (*team_service.ExportSprintResponse, error)
		ListSprints(c *gin.Context, req *team_service.ListSprintsRequest) (*team_service.ListSprintsResponse, error)
		UpdateSprint(c *gin.Context, req *team_service.UpdateSprintRequest) (*team_service.UpdateSprintResponse, error)
		UpdateSprintStatus(c *gin.Context, req *team_service.UpdateSprintStatusRequest) (*team_service.UpdateSprintStatusResponse, error)
		DeleteSprint(c *gin.Context, req *common.IDRequest) (*team_service.DeleteSprintResponse, error)
	}
	WorkClient interface {
		CreateWork(c *gin.Context, req *team_service.CreateWorkRequest) (*team_service.CreateWorkResponse, error)
		GetWork(c *gin.Context, req *common.IDRequest) (*team_service.GetWorkResponse, error)
		ListWorks(c *gin.Context, req *team_service.ListWorksRequest) (*team_service.ListWorksResponse, error)
		UpdateWork(c *gin.Context, req *team_service.UpdateWorkRequest) (*team_service.UpdateWorkResponse, error)
		DeleteWork(c *gin.Context, req *common.IDRequest) (*team_service.DeleteWorkResponse, error)

		CreateChecklistItem(c *gin.Context, req *team_service.CreateChecklistItemRequest) (*team_service.CreateChecklistItemResponse, error)
		UpdateChecklistItem(c *gin.Context, req *team_service.UpdateChecklistItemRequest) (*team_service.UpdateChecklistItemResponse, error)
		DeleteChecklistItem(c *gin.Context, req *common.IDRequest) (*team_service.DeleteChecklistItemResponse, error)

		CreateComment(c *gin.Context, req *team_service.CreateCommentRequest) (*team_service.CreateCommentResponse, error)
		UpdateComment(c *gin.Context, req *team_service.UpdateCommentRequest) (*team_service.UpdateCommentResponse, error)
		DeleteComment(c *gin.Context, req *common.IDRequest) (*team_service.DeleteCommentResponse, error)
	}
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
