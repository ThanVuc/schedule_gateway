package user_route

import (
	"schedule_gateway/internal/controller"
	v1 "schedule_gateway/internal/grpc/auth.v1"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (p *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// wire the controller
	userController := controller.NewUserController()
	// private router
	permissionRouterPrivate := Router.Group("user-info")
	{
		permissionRouterPrivate.GET("/:id", middlewares.CheckPerm("user-info", "readOne"), userController.GetUserInfo)

		permissionRouterPrivate.PUT("/:id/update", middlewares.CheckPerm("user-info", "update"), userController.UpdateUserInfo)
	}
	RegisterPermissionRouterResource()
}

func RegisterPermissionRouterResource() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.UserResource.ResourceId)
	register.AddResource(resoucePredefine.UserResource, []*v1.Action{
		{
			ActionId: register.GenerateActionId(),
			Action:   "readOne",
		},
		{
			ActionId: register.GenerateActionId(),
			Action:   "update",
		},
	})
}
