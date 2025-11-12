package personalschedule_router

import (
	controller "schedule_gateway/internal/controller/personalschedule"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type GoalRouter struct{}

func (r *GoalRouter) InitGoalRouter(Router *gin.RouterGroup) {
	// wire the controller
	goalController := controller.NewGoalController()
	// private router
	goalRouterPrivate := Router.Group("goals")
	{
		goalRouterPrivate.GET("", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.READ_GOALS_ACTION), goalController.GetGoals)
		goalRouterPrivate.GET("/:id", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.READ_ONE_GOAL_ACTION), goalController.GetGoal)
		goalRouterPrivate.POST("", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.CREATE_GOAL_ACTION), goalController.UpsertGoal)
		goalRouterPrivate.POST("/:id", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.UPDATE_GOAL_ACTION), goalController.UpsertGoal)
		goalRouterPrivate.DELETE("/:id", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.DELETE_GOAL_ACTION), goalController.DeleteGoal)
	}
	RegisterGoalRouterResouce()

}

func RegisterGoalRouterResouce() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.GoalResource.Id)

	register.AddResource(resoucePredefine.GoalResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_GOALS_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_ONE_GOAL_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.CREATE_GOAL_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.UPDATE_GOAL_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.DELETE_GOAL_ACTION,
		},
	})
}
