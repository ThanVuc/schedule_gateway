package personalschedule_router

import (
	controller "schedule_gateway/internal/controller/personalschedule"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type WorkRouter struct{}

func (r WorkRouter) InitWorkRouter(Router *gin.RouterGroup) {
	// wire the controller
	workController := controller.NewWorkController()
	// private router
	workRouterPrivate := Router.Group("works")
	{
		workRouterPrivate.POST("", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.CREATE_ACTION), workController.UpsertWork)
		workRouterPrivate.POST("/:id", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.UPDATE_ACTION), workController.UpsertWork)
		workRouterPrivate.GET("", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.READ_ALL_ACTION), workController.GetWorks)
		workRouterPrivate.GET("/:id", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.READ_ONE_ACTION), workController.GetWork)
		workRouterPrivate.DELETE("/:id", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.DELETE_ACTION), workController.DeleteWork)
		workRouterPrivate.POST("/recovery", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.RECOVER_WORKS_ACTION), workController.GetRecoveryWorks)
		workRouterPrivate.PATCH("/:id", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.UPDATE_ACTION), workController.UpdateWorkLabel)
	}
	RegisterWorkRouterResouce()

}

func RegisterWorkRouterResouce() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.WorkResource.Id)

	register.AddResource(resoucePredefine.WorkResource, []*auth.Action{
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
			Name: constant.READ_ALL_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_ONE_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.DELETE_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.RECOVER_WORKS_ACTION,
		},
	})
}
