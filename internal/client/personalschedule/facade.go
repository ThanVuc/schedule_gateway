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
