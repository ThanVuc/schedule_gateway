package authorization

import (
	"schedule_gateway/internal/controller"
	v1 "schedule_gateway/internal/grpc/auth.v1"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	// wire the controller
	roleController := controller.NewRoleController()
	// private router
	roleRouterPrivate := Router.Group("/roles")
	{
		roleRouterPrivate.GET("/", middlewares.CheckPerm("roles", "readAll"), roleController.GetRoles)

		roleRouterPrivate.POST("/", middlewares.CheckPerm("roles", "create"), roleController.CreateRole)

		roleRouterPrivate.PUT("/:id", middlewares.CheckPerm("roles", "update"), roleController.UpdateRole)

		roleRouterPrivate.DELETE("/:id", middlewares.CheckPerm("roles", "delete"), roleController.DeleteRole)

		roleRouterPrivate.PUT("/:id/disable-or-enable", middlewares.CheckPerm("roles", "disableOrEnable"), roleController.DisableOrEnableRole)

		roleRouterPrivate.POST("/assign-role-to-user", middlewares.CheckPerm("roles", "assignToUser"), roleController.AssignRoleToUser)
	}
	RegisterRoleRouterResouce()
}

func RegisterRoleRouterResouce() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.RoleResource.ResourceId)

	register.AddResource(resoucePredefine.RoleResource, []*v1.Action{
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
			Action:   "disableOrEnable",
		},
		{
			ActionId: register.GenerateActionId(),
			Action:   "assignToUser",
		},
	})
}
