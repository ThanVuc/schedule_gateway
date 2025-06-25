package authentication

import (
	"schedule_gateway/internal/controller"
	v1 "schedule_gateway/internal/grpc/auth.v1"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (ar *AuthRouter) InitAuthRouter(routerGroup *gin.RouterGroup) {
	authController := controller.NewAuthController()

	authRouterPrivate := routerGroup.Group("auth")
	{
		authRouterPrivate.POST("logout", middlewares.CheckPerm("auth", "logout"), authController.Logout)
		authRouterPrivate.POST("reset-password", middlewares.CheckPerm("auth", "reset"), authController.ResetPassword)
		authRouterPrivate.POST("forgot-password", middlewares.CheckPerm("auth", "retrieve"), authController.ForgotPassword)
	}

	authRouterPublic := routerGroup.Group("auth")
	{
		authRouterPublic.POST("register", authController.Register)
		authRouterPublic.POST("login", authController.Login)
		authRouterPrivate.POST("confirm-email", authController.ConfirmEmail)
		authRouterPublic.POST("confirm-forgot-password", authController.ConfirmForgotPassword)
	}

	RegisterAuthRouterResource()
}

func RegisterAuthRouterResource() {
	// Register the resources and their permissions
	resoucePredefine := helper.InitResources()
	register := helper.NewResourceRegiseter(resoucePredefine.AuthResource.ResourceId)

	register.AddResource(resoucePredefine.AuthResource, []*v1.Action{
		{
			ActionId: register.GenerateActionId(),
			Action:   "logout",
		},
		{
			ActionId: register.GenerateActionId(),
			Action:   "reset",
		},
		{
			ActionId: register.GenerateActionId(),
			Action:   "retrieve",
		},
	})
}
