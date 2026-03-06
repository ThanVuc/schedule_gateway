package team_router

import (
	controller "schedule_gateway/internal/controller/team"
	"schedule_gateway/internal/helper"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type WorkRouter struct{}

func (r *WorkRouter) InitWorkRouter(Router *gin.RouterGroup) {
	workController := controller.NewWorkController()
	// private router
	workRouterPrivate := Router.Group("ts/works")
	{
		// TODO: add permission
		println(workController)
		println(workRouterPrivate)
	}
	r.Register()
}

func (r *WorkRouter) Register() {
	// Register the resources and their works
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.TeamWorkResource.Id)
	register.AddResource(resoucePredefine.TeamWorkResource, []*auth.Action{})
}
