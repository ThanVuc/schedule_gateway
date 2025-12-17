package notification_router

import (
	controller "schedule_gateway/internal/controller/notification"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type NotificationRouter struct{}

func (r *NotificationRouter) InitNotificationRouter(Router *gin.RouterGroup) {
	// wire the controller
	notificationController := controller.NewNotificationController()
	userUserNotificationController := controller.NewUserNotificationController()
	// private router
	notifcationRouterPrivate := Router.Group("notifications")
	{
		notifcationRouterPrivate.GET("", middlewares.CheckPerm(constant.NOTIFICATION_RESOURCE, constant.READ_ALL_ACTION), notificationController.GetNotificationsByRecipientId)
		notifcationRouterPrivate.POST("/upsert-fcm", middlewares.CheckPerm(constant.NOTIFICATION_RESOURCE, constant.SAVE_FCM_TOKEN_ACTION), userUserNotificationController.UpsertUserFCMToken)
		notifcationRouterPrivate.POST("/mark-as-read", middlewares.CheckPerm(constant.NOTIFICATION_RESOURCE, constant.MARK_AS_READ_ACTION), notificationController.MarkNotificationsAsRead)
		notifcationRouterPrivate.DELETE(":id", middlewares.CheckPerm(constant.NOTIFICATION_RESOURCE, constant.DELETE_ACTION), notificationController.DeleteNotificationById)
	}
	RegisterNotificationRouterResouce()
}

func RegisterNotificationRouterResouce() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.NotificationResource.Id)

	register.AddResource(resoucePredefine.NotificationResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: constant.SAVE_FCM_TOKEN_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.MARK_AS_READ_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_ALL_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.DELETE_ACTION,
		},
	})
}
