package authentication

import (
	"schedule_gateway/internal/controller"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	"schedule_gateway/proto/auth"

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
	register := helper.NewResourceRegiseter(resoucePredefine.AuthResource.Id)

	register.AddResource(resoucePredefine.AuthResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: "logout",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "reset",
		},
		{
			Id:   register.GenerateActionId(),
			Name: "retrieve",
		},
	})
}
