package authentication

import (
	"net/http"
	"schedule_gateway/global"
	controller "schedule_gateway/internal/controller/auth"

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
}
