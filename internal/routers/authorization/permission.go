package authorization

import (
	"schedule_gateway/internal/controller"
	v1 "schedule_gateway/internal/grpc/auth.v1"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type PermissionRouter struct{}

func (p *PermissionRouter) InitPermissionRouter(Router *gin.RouterGroup) {
	// wire the controller
	permissionController := controller.NewPermissionController()
	// private router
	permissionRouterPrivate := Router.Group("permissions")
	{
		permissionRouterPrivate.GET("/", middlewares.CheckPerm("permissions", "readAll"), permissionController.GetPermissions)

		permissionRouterPrivate.POST("/", middlewares.CheckPerm("permissions", "create"), permissionController.CreatePermission)

		permissionRouterPrivate.PUT("/:id", middlewares.CheckPerm("permissions", "update"), permissionController.UpdatePermission)

		permissionRouterPrivate.DELETE("/:id", middlewares.CheckPerm("permissions", "delete"), permissionController.DeletePermission)

		permissionRouterPrivate.POST("/assign-permission-to-role", middlewares.CheckPerm("permissions", "assign"), permissionController.AssignPermissionToRole)
	}
	RegisterPermissionRouterResource()
}

func RegisterPermissionRouterResource() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.PermissionResource.ResourceId)
	register.AddResource(resoucePredefine.PermissionResource, []*v1.Action{
		{
			ActionId: register.GenerateActionId(),
			Action:   "readAll",
		},
		{
			ActionId: register.GenerateActionId(),
			Action:   "create",
		},
		{
			ActionId: register.GenerateActionId(),
			Action:   "update",
		},
		{
			ActionId: register.GenerateActionId(),
			Action:   "delete",
		},
		{
			ActionId: register.GenerateActionId(),
			Action:   "assign",
		},
	})
}
