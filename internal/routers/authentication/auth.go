package authentication

import (
	controller "schedule_gateway/internal/controller/auth"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
}

func (ar *AuthRouter) InitAuthRouter(routerGroup *gin.RouterGroup) {
	authController := controller.NewAuthController()

	authRouterPublic := routerGroup.Group("auth")
	{
		authRouterPublic.POST("login-with-google", authController.LoginWithGoogle)
	}
}
