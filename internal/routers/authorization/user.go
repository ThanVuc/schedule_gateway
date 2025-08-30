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
		userRouterPrivate.GET("", middlewares.CheckPerm(constant.ADMIN_USER_RESOURCE, constant.READ_ALL_USERS_ACTION), userController.GetUsers)
		userRouterPrivate.GET("/:id", middlewares.CheckPerm(constant.ADMIN_USER_RESOURCE, constant.READ_ONE_USER_ACTION), userController.GetUser)
		userRouterPrivate.PUT("lock-user", middlewares.CheckPerm(constant.ADMIN_USER_RESOURCE, constant.LOCK_USER_ACTION),  userController.LockOrUnLockUser)
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
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_ALL_USERS_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_ONE_USER_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.LOCK_USER_ACTION,
		},
	})
}
