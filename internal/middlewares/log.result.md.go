package middlewares

import (
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func LogResultMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process the request
		start := time.Now()
		c.Next()

		status := c.Writer.Status()
		latency := time.Since(start).Milliseconds()
		method := c.Request.Method
		path := c.Request.URL.Path
		requestId, _ := c.Get("requestId")

		logString := fmt.Sprintf("\033[34m'path': %s, 'status': %d, 'latency': %dms, 'method': %s, 'requestId': %s\033[0m",
			path, status, latency, method, requestId)
		log.Println(logString)
	}
}
