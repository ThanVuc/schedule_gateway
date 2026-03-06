package team_router

import (
	controller "schedule_gateway/internal/controller/team"
	"schedule_gateway/internal/helper"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type SprintRouter struct{}

func (r *SprintRouter) InitSprintRouter(Router *gin.RouterGroup) {
	sprintController := controller.NewSprintController()
	// private router
	sprintRouterPrivate := Router.Group("ts/sprints")
	{
		// TODO: add permission
		println(sprintController)
		println(sprintRouterPrivate)
	}
	r.Register()
}

func (r *SprintRouter) Register() {
	// Register the resources and their sprints
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.SprintResource.Id)
	register.AddResource(resoucePredefine.SprintResource, []*auth.Action{})
}
