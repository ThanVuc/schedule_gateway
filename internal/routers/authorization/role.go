package authorization

import (
	"schedule_gateway/internal/controller"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) {
	// wire the controller
	roleController := controller.NewRoleController()
	// private router
	roleRouterPrivate := Router.Group("roles")
	{
		roleRouterPrivate.GET("", middlewares.CheckPerm(constant.ROLE_RESOURCE, constant.READ_ALL_ACTION), roleController.GetRoles)

		roleRouterPrivate.GET("/:id", middlewares.CheckPerm(constant.ROLE_RESOURCE, constant.READ_ONE_ACTION), roleController.GetRole)

		roleRouterPrivate.DELETE("/:id", middlewares.CheckPerm(constant.ROLE_RESOURCE, constant.DELETE_ACTION), roleController.DeleteRole)

		roleRouterPrivate.PUT("/:id/disable-or-enable", middlewares.CheckPerm(constant.ROLE_RESOURCE, constant.ENABLE_AND_DISABLE_ACTION), roleController.DisableOrEnableRole)

		roleRouterPrivate.POST("", middlewares.CheckPerm(constant.ROLE_RESOURCE, constant.CREATE_ACTION), roleController.UpsertRole)

		roleRouterPrivate.PUT("/:id", middlewares.CheckPerm(constant.ROLE_RESOURCE, constant.UPDATE_ACTION), roleController.UpsertRole)

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
			Name: constant.ENABLE_AND_DISABLE_ACTION,
		},
	})
}
