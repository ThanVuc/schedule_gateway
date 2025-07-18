package authentication

import (
	"schedule_gateway/internal/controller"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	"schedule_gateway/proto/auth"

	"github.com/gin-gonic/gin"
)

type TokenRouter struct{}

func (tr *TokenRouter) InitTokenRouter(routerGroup *gin.RouterGroup) {
	tokenController := controller.NewTokenController()

	tokenRouterPrivate := routerGroup.Group("token")
	{
		tokenRouterPrivate.POST("revoke", middlewares.CheckPerm("token", "revoke"), tokenController.RevokeToken)
	}

	tokenRouterPublic := routerGroup.Group("token")
	{
		tokenRouterPublic.POST("refresh", tokenController.RefreshToken)
	}

	RegisterTokenRouterResource()
}

func RegisterTokenRouterResource() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()
	register := helper.NewResourceRegiseter(resoucePredefine.TokenResource.Id)

	register.AddResource(resoucePredefine.TokenResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: "revoke",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "refresh",
		},
	})
}
