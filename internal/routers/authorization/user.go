package authorization

import (
	controller "schedule_gateway/internal/controller/auth"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userController := controller.NewUserController()

	userRouterPrivate := Router.Group("users")
	{
		userRouterPrivate.POST("assign-role", middlewares.CheckPerm(constant.ADMIN_USER_RESOURCE, constant.ASSIGN_ROLE_ACTION), userController.AssignRoleToUser)
	}
	RegisterUserRouterResouce()
}

func RegisterUserRouterResouce() {
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.AdminUserResource.Id)

	register.AddResource(resoucePredefine.AdminUserResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: constant.ASSIGN_ROLE_ACTION,
		},
	})
}
