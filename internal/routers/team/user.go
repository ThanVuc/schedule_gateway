package team_router

import (
	controller "schedule_gateway/internal/controller/team"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController := controller.NewUserController()

	userRouter := Router.Group("ts/users")
	{
		userRouter.GET("",
			middlewares.CheckPerm(constant.TEAM_USER_RESOURCE, constant.READ_ONE_ACTION),
			userController.GetUserInfo,
		)
		userRouter.PATCH("/configurations",
			middlewares.CheckPerm(constant.TEAM_USER_RESOURCE, constant.UPDATE_NOTIFICATION_CONFIGURATION_ACTION),
			userController.NotificationConfiguration,
		)

	}

	u.Register()
}

func (u *UserRouter) Register() {
	// Register the resources and their sprints
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.TeamUserResource.Id)
	register.AddResource(resoucePredefine.TeamUserResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_ONE_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.UPDATE_NOTIFICATION_CONFIGURATION_ACTION,
		},
	})
}
