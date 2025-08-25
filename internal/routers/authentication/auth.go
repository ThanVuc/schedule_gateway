package authentication

import (
	"net/http"
	"schedule_gateway/global"
	controller "schedule_gateway/internal/controller/auth"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	constant "schedule_gateway/internal/routers/constant"
	"schedule_gateway/proto/auth"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
	"go.uber.org/zap"
)

type AuthRouter struct {
}

func (ar *AuthRouter) InitAuthRouter(routerGroup *gin.RouterGroup) {
	authController := controller.NewAuthController()

	authRouterPublic := routerGroup.Group("auth")
	{
		authRouterPublic.POST("login-with-google", authController.LoginWithGoogle)
		authRouterPublic.GET("csrf-token", func(ctx *gin.Context) {
			session := sessions.Default(ctx)
			session.Set("init", true)
			if err := session.Save(); err != nil {
				global.Logger.Error("Failed to save session", "", zap.Error(err))
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
				return
			}

			ctx.JSON(200, gin.H{
				"csrf_token": csrf.GetToken(ctx),
			})
		})
	}

	authRouterPrivate := routerGroup.Group("auth")
	{
		authRouterPrivate.POST("refresh-token", middlewares.CheckPerm(constant.AUTH_RESOURCE, constant.REFRESH_TOKEN_ACTION), authController.RefreshToken)
		authRouterPrivate.POST("logout", middlewares.CheckPerm(constant.AUTH_RESOURCE, constant.LOGOUT_ACTION), authController.Logout)
	}

	RegisterAuthRouterResource()
}

func RegisterAuthRouterResource() {
	resoucePredefine := helper.InitResources()

	register := helper.NewResourceRegiseter(resoucePredefine.AuthResource.Id)
	register.AddResource(resoucePredefine.AuthResource, []*auth.Action{
		{
			Id:   register.GenerateActionId(),
			Name: constant.REFRESH_TOKEN_ACTION,
		},
		{
			Id:   register.GenerateActionId(),
			Name: constant.LOGOUT_ACTION,
		},
	})
}
