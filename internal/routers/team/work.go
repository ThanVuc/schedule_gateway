package team_router

import (
	controller "schedule_gateway/internal/controller/team"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type WorkRouter struct{}

func (r *WorkRouter) InitWorkRouter(Router *gin.RouterGroup) {
	workController := controller.NewWorkController()

	workRouter := Router.Group("ts/groups/:group_id/works")
	{
		workRouter.POST("",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.CREATE_ACTION),
			workController.CreateWork,
		)

		workRouter.GET("",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.READ_ALL_ACTION),
			workController.ListWorks,
		)

		workRouter.GET("/:work_id",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.READ_ONE_ACTION),
			workController.GetWork,
		)

		workRouter.PATCH("/:work_id",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.UPDATE_ACTION),
			workController.UpdateWork,
		)

		workRouter.DELETE("/:work_id",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.DELETE_ACTION),
			workController.DeleteWork,
		)

		workRouter.POST("/:work_id/checklists",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.CREATE_ACTION),
			workController.CreateChecklistItem,
		)

		workRouter.PATCH("/:work_id/checklists/:checklist_id",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.UPDATE_ACTION),
			workController.UpdateChecklistItem,
		)

		workRouter.DELETE("/:work_id/checklists/:checklist_id",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.DELETE_ACTION),
			workController.DeleteChecklistItem,
		)

		workRouter.POST("/:work_id/comments",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.CREATE_ACTION),
			workController.CreateComment,
		)

		workRouter.PATCH("/:work_id/comments/:comment_id",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.UPDATE_ACTION),
			workController.UpdateComment,
		)

		workRouter.DELETE("/:work_id/comments/:comment_id",
			middlewares.CheckPerm(constant.TEAM_WORK_RESOURCE, constant.DELETE_ACTION),
			workController.DeleteComment,
		)
	}

	r.Register()
}

func (r *WorkRouter) Register() {
	// Register the resources and their works
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.TeamWorkResource.Id)
	register.AddResource(resoucePredefine.TeamWorkResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: constant.CREATE_ACTION,
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
			Name: constant.UPDATE_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.DELETE_ACTION,
		},
	})
}
