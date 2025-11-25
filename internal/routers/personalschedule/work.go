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
		workRouterPrivate.POST("", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.CREATE_WORK_ACTION), workController.UpsertWork)
		workRouterPrivate.POST("/:id", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.UPDATE_WORK_ACTION), workController.UpsertWork)
		workRouterPrivate.GET("", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.READ_WORKS_ACTION), workController.GetWorks)
		workRouterPrivate.GET("/:id", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.READ_WORK_ACTION), workController.GetWork)
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
			Name: constant.CREATE_WORK_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.UPDATE_WORK_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_WORKS_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_WORK_ACTION,
		},
	})
}
