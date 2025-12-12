package notification_client

import (
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/pkg/settings"
	"schedule_gateway/proto/common"
	"schedule_gateway/proto/notification_service"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type (
	NotificationClient interface {
		GetNotificationsByRecipientId(c *gin.Context, req *common.IDRequest) (*notification_service.GetNotificationsByRecipientIdResponse, error)
		MarkNotificationAsRead(c *gin.Context, req *common.IDsRequest) (*common.EmptyResponse, error)
		DeleteNotificationById(c *gin.Context, req *common.IDRequest) (*common.EmptyResponse, error)
	}

	UserNotificationClient interface {
		UpsertUserFCMToken(c *gin.Context, req *notification_service.UpsertUserFCMTokenRequest) (*common.EmptyResponse, error)
	}
)

func getConn(baseConfig settings.GrpcBase) *grpc.ClientConn {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", baseConfig.GetHost(), baseConfig.GetPort()), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic("Failed to connect to gRPC server: " + err.Error())
	}
	return conn
}

func NewNotificationClient() NotificationClient {
	conn := getConn(&global.Config.NotificationService)

	client := notification_service.NewNotificationServiceClient(conn)
	if client == nil {
		panic("Failed to create NotificationService client at " + fmt.Sprintf("%s:%d", global.Config.PersonalScheduleService.GetHost(), global.Config.PersonalScheduleService.GetPort()))
	}

	return &notificationClient{
		logger:             global.Logger,
		notificationClient: client,
	}
}

func NewUserNotificationClient() UserNotificationClient {
	conn := getConn(&global.Config.NotificationService)

	client := notification_service.NewUserNotificationServiceClient(conn)
	if client == nil {
		panic("Failed to create NotificationService client at " + fmt.Sprintf("%s:%d", global.Config.PersonalScheduleService.GetHost(), global.Config.PersonalScheduleService.GetPort()))
	}

	return &userNotificationClient{
		logger:                 global.Logger,
		userNotificationClient: client,
	}
}
