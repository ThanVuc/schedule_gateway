package authentication

import (
	"schedule_gateway/internal/controller"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"

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
	helper.AddResource("token", []string{
		"refresh",
		"revoke",
	})
}
