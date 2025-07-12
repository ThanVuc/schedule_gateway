package user_route

import (
	"schedule_gateway/internal/controller"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	"schedule_gateway/proto/auth"

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

	register := helper.NewResourceRegiseter(resoucePredefine.UserResource.Id)
	register.AddResource(resoucePredefine.UserResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: "readOne",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "update",
		},
	})
}
