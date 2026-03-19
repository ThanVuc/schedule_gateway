package team_router

import (
	controller "schedule_gateway/internal/controller/team"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
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
		groupRouterPrivate.GET("/ping", middlewares.CheckPerm(constant.WORK_RESOURCE, constant.CREATE_ACTION), groupController.Ping)
		groupRouterPrivate.POST("", middlewares.CheckPerm(constant.GROUP_RESOURCE, constant.CREATE_ACTION), groupController.CreateGroup)
		groupRouterPrivate.GET("/:group_id", middlewares.CheckPerm(constant.GROUP_RESOURCE, constant.READ_ONE_ACTION), groupController.GetGroup)
		groupRouterPrivate.PATCH("/:group_id", middlewares.CheckPerm(constant.GROUP_RESOURCE, constant.UPDATE_ACTION), groupController.UpdateGroup)
		groupRouterPrivate.DELETE("/:group_id", middlewares.CheckPerm(constant.GROUP_RESOURCE, constant.DELETE_ACTION), groupController.DeleteGroup)
		groupRouterPrivate.GET("/:group_id/members", middlewares.CheckPerm(constant.GROUP_RESOURCE, constant.READ_LIST_MEMBERS_ACTION), groupController.ListMembers)
		groupRouterPrivate.PATCH("/:group_id/members/:user_id", middlewares.CheckPerm(constant.GROUP_RESOURCE, constant.UPDATE_MEMBER_ROLE_ACTION), groupController.UpdateMemberRole)
		groupRouterPrivate.DELETE("/:group_id/members/:user_id", middlewares.CheckPerm(constant.GROUP_RESOURCE, constant.REMOVE_MEMBER_ACTION), groupController.RemoveMember)
		groupRouterPrivate.POST("/:group_id/invites", middlewares.CheckPerm(constant.GROUP_RESOURCE, constant.CREATE_INVITE_ACTION), groupController.CreateInvite)
	}
	r.Register()
}

func (r *GroupRouter) Register() {
	// Register the resources and their groups
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.GroupResource.Id)
	register.AddResource(resoucePredefine.GroupResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: constant.CREATE_ACTION,
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
		{
			Id:   register.GenerateActionId(),
			Name: constant.READ_LIST_MEMBERS_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.UPDATE_MEMBER_ROLE_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.REMOVE_MEMBER_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.CREATE_INVITE_ACTION,
		},
	})
}
