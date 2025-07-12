package initialize

import (
	"schedule_gateway/internal/middlewares"
	"schedule_gateway/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middlewares.TrackLogMiddleware())

	permRouter := routers.RouterGroupApp.AuthorizationRouterEnter.PermissionRouter
	roleRouter := routers.RouterGroupApp.AuthorizationRouterEnter.RoleRouter
	tokenRouter := routers.RouterGroupApp.AuthenticationRouterEnter.TokenRouter
	authRouter := routers.RouterGroupApp.AuthenticationRouterEnter.AuthRouter
	userRouter := routers.RouterGroupApp.UserRouterEnter.UserRouter

	MainGroup := r.Group("api/v1/")
	{
		MainGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Gateway is running",
			})
		})
	}
	{
		permRouter.InitPermissionRouter(MainGroup)
		roleRouter.InitRoleRouter(MainGroup)
		tokenRouter.InitTokenRouter(MainGroup)
		authRouter.InitAuthRouter(MainGroup)
		userRouter.InitUserRouter(MainGroup)
	}
}
