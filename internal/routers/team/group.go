package team_router

import (
	controller "schedule_gateway/internal/controller/team"
	"schedule_gateway/internal/helper"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type GroupRouter struct{}

func (r *GroupRouter) InitGroupRouter(Router *gin.RouterGroup) {
	groupController := controller.NewGroupController()
	// private router
	groupRouterPrivate := Router.Group("ts/groups")
	{
		// TODO: add permission
		println(groupController)
		println(groupRouterPrivate)
	}
	r.Register()
}

func (r *GroupRouter) Register() {
	// Register the resources and their groups
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.GroupResource.Id)
	register.AddResource(resoucePredefine.GroupResource, []*auth.Action{})
}
