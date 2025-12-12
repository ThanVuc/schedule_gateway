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
		goalRouterPrivate.GET("", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.READ_ALL_ACTION), goalController.GetGoals)
		goalRouterPrivate.GET("/:id", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.READ_ONE_ACTION), goalController.GetGoal)
		goalRouterPrivate.POST("", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.CREATE_ACTION), goalController.UpsertGoal)
		goalRouterPrivate.POST("/:id", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.UPDATE_ACTION), goalController.UpsertGoal)
		goalRouterPrivate.DELETE("/:id", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.DELETE_ACTION), goalController.DeleteGoal)
		goalRouterPrivate.GET("/dialog/list", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.READ_GOALS_FOR_DIALOG_ACTION), goalController.GetGoalsForDialog)
		goalRouterPrivate.PATCH("/:id", middlewares.CheckPerm(constant.GOAL_RESOURCE, constant.UPDATE_ACTION), goalController.UpdateGoalLabel)
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
			Name: constant.READ_GOALS_FOR_DIALOG_ACTION,
		},
	})
}
