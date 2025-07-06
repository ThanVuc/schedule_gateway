package authorization

import (
	"schedule_gateway/internal/controller"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	"schedule_gateway/proto/auth"

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

		roleRouterPrivate.DELETE("/:id", middlewares.CheckPerm("roles", "delete"), roleController.DeleteRole)

		roleRouterPrivate.PUT("/:id/disable-or-enable", middlewares.CheckPerm("roles", "disableOrEnable"), roleController.DisableOrEnableRole)

	}
	RegisterRoleRouterResouce()
}

func RegisterRoleRouterResouce() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.RoleResource.Id)

	register.AddResource(resoucePredefine.RoleResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: "readAll",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "delete",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "disableOrEnable",
		},
	})
}
