package authorization

import (
	"schedule_gateway/internal/controller"
	"schedule_gateway/internal/grpc/auth"
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
			Name: "assign",
		},
	})
}
