package user_route

import (
	controller "schedule_gateway/internal/controller/user"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (p *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController := controller.NewUserController()
	// private router
	permissionRouterPrivate := Router.Group("profile")
	{
		permissionRouterPrivate.GET("", middlewares.CheckPerm(constant.USER_RESOURCE, constant.READ_ONE_ACTION), userController.GetUserProfile)
		permissionRouterPrivate.PUT("/update", middlewares.CheckPerm(constant.USER_RESOURCE, constant.UPDATE_ACTION), userController.UpdateUserInfo)
	}
	RegisterPermissionRouterResource()
}

func RegisterPermissionRouterResource() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.UserResource.Id)
	register.AddResource(resoucePredefine.UserResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_ONE_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.UPDATE_ACTION,
		},
	})
}
