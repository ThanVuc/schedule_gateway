package initialize

import (
	"fmt"
	"schedule_gateway/global"
	"schedule_gateway/internal/helper"
	"schedule_gateway/internal/middlewares"

	"github.com/gin-gonic/gin"
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

	// producer := producer.NewProducer()

	// init the app with gin
	// This order is important, as the middleware needs to be set before the routes are initialized.
	var r *gin.Engine = gin.New()
	r.Use(middlewares.LogResultMiddleware())
	r.Use(middlewares.TrackLogMiddleware())
	r.Use(middlewares.ErrorHandler())
	InitRouter(r)

	helper.WriteToJsonFile("resources")
	// producer.SendResources()
	go InitResource()

	r.Run(fmt.Sprintf("%s:%d", global.Config.Server.Host, global.Config.Server.Port)) // listen and serve on host:port
}
