package initialize

import (
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"
	"schedule_gateway/pkg/response"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	csrf "github.com/utrack/gin-csrf"
)

/*
@Author: Sinh
@Date: 2025/6/1
@Description: Run initializes the application by loading the configuration,
establishing database connections, and setting up the HTTP server with the specified routes.
@Note: This function is the entry point for the application, setting up the necessary components
*/
func Run() {
	LoadConfig()
	InitLogger()

	var r *gin.Engine = gin.New()
	r.Use(middlewares.LogResultMiddleware())
	r.Use(middlewares.TrackLogMiddleware())
	r.Use(middlewares.CORSMiddleware())

	store := cookie.NewStore([]byte(global.Config.SessionSecret))
	store.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		MaxAge:   24 * 60 * 60,
	})
	r.Use(sessions.Sessions("schdulr_session", store))
	r.Use(csrf.Middleware(csrf.Options{
		Secret:        global.Config.CsrfSecret,
		IgnoreMethods: []string{"GET", "HEAD", "OPTIONS"},
		ErrorFunc: func(c *gin.Context) {
			response.Forbidden(c, "CSRF token mismatch")
			c.Abort()
		},
	}))
	r.Use(middlewares.ErrorHandler())

	InitRouter(r)

	helper.WriteToJsonFile("resources")
	go InitResource()

	r.Run(fmt.Sprintf("%s:%d", global.Config.Server.Host, global.Config.Server.Port))
}
