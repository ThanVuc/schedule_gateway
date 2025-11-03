package initialize

import (
	"schedule_gateway/internal/middlewares"
	"schedule_gateway/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middlewares.TrackLogMiddleware())
	r.Use(middlewares.CORSMiddleware())

	permRouter := routers.RouterGroupApp.AuthorizationRouterEnter.PermissionRouter
	roleRouter := routers.RouterGroupApp.AuthorizationRouterEnter.RoleRouter
	authRouter := routers.RouterGroupApp.AuthenticationRouterEnter.AuthRouter
	authUserRouter := routers.RouterGroupApp.AuthorizationRouterEnter.UserRouter
	userRouter := routers.RouterGroupApp.UserRouterEnter.UserRouter
	labelRouter := routers.RouterGroupApp.PersonalScheduleRouterEnter.LabelRouter
	goalRouter := routers.RouterGroupApp.PersonalScheduleRouterEnter.GoalRouter

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
		authRouter.InitAuthRouter(MainGroup)
		userRouter.InitUserRouter(MainGroup)
		authUserRouter.InitUserRouter(MainGroup)
		labelRouter.InitLabelRouter(MainGroup)
		goalRouter.InitGoalRouter(MainGroup)
	}
}
