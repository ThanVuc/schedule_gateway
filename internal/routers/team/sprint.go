package team_router

import (
	controller "schedule_gateway/internal/controller/team"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type SprintRouter struct{}

func (r *SprintRouter) InitSprintRouter(Router *gin.RouterGroup) {
	sprintController := controller.NewSprintController()

	sprintRouter := Router.Group("ts/groups/:group_id/sprints")
	{
		sprintRouter.POST("",
			middlewares.CheckPerm(constant.SPRINT_RESOURCE, constant.CREATE_ACTION),
			sprintController.CreateSprint,
		)

		sprintRouter.GET("",
			middlewares.CheckPerm(constant.SPRINT_RESOURCE, constant.READ_ALL_ACTION),
			sprintController.ListSprints,
		)

		sprintRouter.GET("/:sprint_id",
			middlewares.CheckPerm(constant.SPRINT_RESOURCE, constant.READ_ONE_ACTION),
			sprintController.GetSprint,
		)

		sprintRouter.PATCH("/:sprint_id",
			middlewares.CheckPerm(constant.SPRINT_RESOURCE, constant.UPDATE_ACTION),
			sprintController.UpdateSprint,
		)

		sprintRouter.PATCH("/:sprint_id/status",
			middlewares.CheckPerm(constant.SPRINT_RESOURCE, constant.UPDATE_ACTION),
			sprintController.UpdateSprintStatus,
		)

		sprintRouter.DELETE("/:sprint_id",
			middlewares.CheckPerm(constant.SPRINT_RESOURCE, constant.DELETE_ACTION),
			sprintController.DeleteSprint,
		)
	}

	r.Register()
}

func (r *SprintRouter) Register() {
	// Register the resources and their sprints
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.SprintResource.Id)
	register.AddResource(resoucePredefine.SprintResource, []*auth.Action{
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
