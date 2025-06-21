package authorization

import (
	"schedule_gateway/internal/controller"
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

		permissionRouterPrivate.POST("/assign-permission-to-role", middlewares.CheckPerm("permissions", "assignToRole"), permissionController.AssignPermissionToRole)
	}
	RegisterPermissionRouterResource()
}

func RegisterPermissionRouterResource() {
	helper.AddResource("permissions", []string{
		"readAll",
		"create",
		"update",
		"delete",
		"assignToRole",
	})
}
