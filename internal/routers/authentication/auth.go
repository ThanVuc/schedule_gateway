package authentication

import (
	"net/http"
	controller "schedule_gateway/internal/controller/auth"

	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

type AuthRouter struct {
}

func (ar *AuthRouter) InitAuthRouter(routerGroup *gin.RouterGroup) {
	authController := controller.NewAuthController()

	authRouterPublic := routerGroup.Group("auth")
	{
		authRouterPublic.POST("login-with-google", authController.LoginWithGoogle)
		authRouterPublic.GET("csrf-token", func(ctx *gin.Context) {
			cookie := &http.Cookie{
				Name:     "init",
				Value:    "true",
				Path:     "/",
				HttpOnly: true,
				Secure:   true,
				SameSite: http.SameSiteNoneMode,
			}
			http.SetCookie(ctx.Writer, cookie)
			ctx.JSON(200, gin.H{
				"csrf_token": csrf.GetToken(ctx),
			})
		})
	}
}
