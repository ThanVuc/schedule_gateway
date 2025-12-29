package personalschedule_client

import (
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/pkg/settings"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/personal_schedule"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	LabelClient interface {
		GetLabelPerTypes(c *gin.Context, req *common.EmptyRequest) (*personal_schedule.GetLabelPerTypesResponse, error)
		GetLabelsByTypeIDs(c *gin.Context, req *common.IDRequest) (*personal_schedule.GetLabelsByTypeIDsResponse, error)
		GetDefaultLabel(c *gin.Context, req *common.EmptyRequest) (*personal_schedule.GetDefaultLabelResponse, error)
	}

	GoalClient interface {
		GetGoals(c *gin.Context, req *personal_schedule.GetGoalsRequest) (*personal_schedule.GetGoalsResponse, error)
		UpsertGoals(c *gin.Context, req *personal_schedule.UpsertGoalRequest) (*personal_schedule.UpsertGoalResponse, error)
		GetGoal(c *gin.Context, req *personal_schedule.GetGoalRequest) (*personal_schedule.GetGoalResponse, error)
		DeleteGoal(c *gin.Context, req *personal_schedule.DeleteGoalRequest) (*personal_schedule.DeleteGoalResponse, error)
		GetGoalForDiaglog(c *gin.Context, req *personal_schedule.GetGoalsForDialogRequest) (*personal_schedule.GetGoalForDialogResponse, error)
		UpdateGoalLabel(c *gin.Context, req *personal_schedule.UpdateGoalLabelRequest) (*personal_schedule.UpdateGoalLabelResponse, error)
	}

	WorkClient interface {
		UpsertWork(c *gin.Context, req *personal_schedule.UpsertWorkRequest) (*personal_schedule.UpsertWorkResponse, error)
		GetWorks(c *gin.Context, req *personal_schedule.GetWorksRequest) (*personal_schedule.GetWorksResponse, error)
		GetWork(c *gin.Context, req *personal_schedule.GetWorkRequest) (*personal_schedule.GetWorkResponse, error)
		DeleteWork(c *gin.Context, req *personal_schedule.DeleteWorkRequest) (*personal_schedule.DeleteWorkResponse, error)
		GetRecoveryWorks(c *gin.Context, req *personal_schedule.GetRecoveryWorksRequest) (*personal_schedule.GetRecoveryWorksResponse, error)
		UpdateWorkLabel(c *gin.Context, req *personal_schedule.UpdateWorkLabelRequest) (*personal_schedule.UpdateWorkLabelResponse, error)
		CommitRecoveryDrafts(c *gin.Context, req *personal_schedule.CommitRecoveryDraftsRequest) (*personal_schedule.CommitRecoveryDraftsResponse, error)
	}
)

func getConn(baseConfig settings.GrpcBase) *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", baseConfig.GetHost(), baseConfig.GetPort()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to gRPC server: " + err.Error())
	}
	return conn
}

func NewLabelClient() LabelClient {
	conn := getConn(&global.Config.PersonalScheduleService)

	client := personal_schedule.NewLabelServiceClient(conn)
	if client == nil {
		panic("Failed to create AuthService client at " + fmt.Sprintf("%s:%d", global.Config.PersonalScheduleService.GetHost(), global.Config.PersonalScheduleService.GetPort()))
	}

	return &labelClient{
		logger:      global.Logger,
		labelClient: client,
	}
}

func NewGoalClient() GoalClient {
	conn := getConn(&global.Config.PersonalScheduleService)

	client := personal_schedule.NewGoalServiceClient(conn)
	if client == nil {
		panic("Failed to create GoalService client at " + fmt.Sprintf("%s:%d", global.Config.PersonalScheduleService.GetHost(), global.Config.PersonalScheduleService.GetPort()))
	}

	return &goalClient{
		logger:     global.Logger,
		goalClient: client,
	}
}

func NewWorkClient() WorkClient {
	conn := getConn(&global.Config.PersonalScheduleService)

	client := personal_schedule.NewWorkServiceClient(conn)
	if client == nil {
		panic("Failed to create WorkService client at " + fmt.Sprintf("%s:%d", global.Config.PersonalScheduleService.GetHost(), global.Config.PersonalScheduleService.GetPort()))
	}

	return &workClient{
		logger:     global.Logger,
		workClient: client,
	}
}
