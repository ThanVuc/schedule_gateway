package authorization

import (
	controller "schedule_gateway/internal/controller/auth"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
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
		permissionRouterPrivate.GET("", middlewares.CheckPerm(constant.PERMISSION_RESOURCE, constant.READ_ALL_ACTION), permissionController.GetPermissions)

		permissionRouterPrivate.GET("/:id", middlewares.CheckPerm(constant.PERMISSION_RESOURCE, constant.READ_ONE_ACTION), permissionController.GetPermission)

		permissionRouterPrivate.POST("", middlewares.CheckPerm(constant.PERMISSION_RESOURCE, constant.CREATE_ACTION), permissionController.UpsertPermission)

		permissionRouterPrivate.PUT("/:id", middlewares.CheckPerm(constant.PERMISSION_RESOURCE, constant.UPDATE_ACTION), permissionController.UpsertPermission)

		permissionRouterPrivate.DELETE("/:id", middlewares.CheckPerm(constant.PERMISSION_RESOURCE, constant.DELETE_ACTION), permissionController.DeletePermission)

		permissionRouterPrivate.GET("/resources", middlewares.CheckPerm(constant.PERMISSION_RESOURCE, constant.READ_RESOURCES_ACTION), permissionController.GetResources)

		permissionRouterPrivate.GET("/actions", middlewares.CheckPerm(constant.PERMISSION_RESOURCE, constant.READ_ACTIONS_ACTION), permissionController.GetActions)
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
			Name: constant.READ_ALL_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_ONE_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.CREATE_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.UPDATE_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.DELETE_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_RESOURCES_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_ACTIONS_ACTION,
		},
	})
}
