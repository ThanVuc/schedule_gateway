package personalschedule_router

import (
	controller "schedule_gateway/internal/controller/personalschedule"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type LabelRouter struct{}

func (r *LabelRouter) InitLabelRouter(Router *gin.RouterGroup) {
	// wire the controller
	labelController := controller.NewLabelController()
	// private router
	labelRouterPrivate := Router.Group("labels")
	{
		labelRouterPrivate.GET("label-per-types", middlewares.CheckPerm(constant.LABEL_RESOURCE, constant.READ_ALL_LABEL_PER_TYPES_ACTION), labelController.GetLabelPerTypes)
		labelRouterPrivate.GET("/types/:type_id", middlewares.CheckPerm(constant.LABEL_RESOURCE, constant.READ_LABELS_BY_TYPE_ACTION), labelController.GetLabelsByTypeIDs)
	}
	RegisterLabelRouterResouce()
}

func RegisterLabelRouterResouce() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.LabelResource.Id)

	register.AddResource(resoucePredefine.LabelResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_ALL_LABEL_PER_TYPES_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_LABELS_BY_TYPE_ACTION,
		},
	})
}
