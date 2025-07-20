package authorization

import (
	"schedule_gateway/internal/controller"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type PermissionRouter struct{}

func (p *PermissionRouter) InitPermissionRouter(Router *gin.RouterGroup) {
	// wire the controller
	permissionController := controller.NewPermissionController()
	// private router
	permissionRouterPrivate := Router.Group("permissions")
	{
		permissionRouterPrivate.GET("", middlewares.CheckPerm("permissions", "readAll"), permissionController.GetPermissions)

		permissionRouterPrivate.GET("/:id", middlewares.CheckPerm("permissions", "readOne"), permissionController.GetPermission)

		permissionRouterPrivate.POST("", middlewares.CheckPerm("permissions", "create"), permissionController.UpsertPermission)

		permissionRouterPrivate.PUT("/:id", middlewares.CheckPerm("permissions", "update"), permissionController.UpsertPermission)

		permissionRouterPrivate.DELETE("/:id", middlewares.CheckPerm("permissions", "delete"), permissionController.DeletePermission)

		permissionRouterPrivate.GET("/resources", middlewares.CheckPerm("permissions", "readResources"), permissionController.GetResources)

		permissionRouterPrivate.GET("/actions", middlewares.CheckPerm("permissions", "readActions"), permissionController.GetActions)
	}
	RegisterPermissionRouterResource()
}

func RegisterPermissionRouterResource() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.PermissionResource.Id)
	register.AddResource(resoucePredefine.PermissionResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: "readAll",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "create",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "update",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "delete",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "readResources",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "readActions",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "readOne",
		},
	})
}
